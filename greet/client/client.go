package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dfreilich/grpc-samples/greet/greetpb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const address = "localhost"
const port = "50051"

var dFGreeting = &greetpb.Greeting{
	FirstName: "David",
	LastName:  "Freilich",
}

func main() {
	log.Println("Starting client")
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
	log.Println("Succesfully ran")
}

func run() error {
	certFile := "ssl/ca.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return errors.Wrap(err, "error while loading CA trust certificates")
	}
	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), opts)
	if err != nil {
		return errors.Wrap(err, "could not connect")
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created client")

	// if err := doUnary(c); err != nil {
	// 	return errors.Wrap(err, "failed to do unary RPC call")
	// }

	// return doServerStreaming(c)
	// return doClientStreaming(c)
	// return doBiDiStreaming(c)
	doUnaryWithDeadline(c, 5*time.Second)
	return doUnaryWithDeadline(c, 1*time.Second)
}

func doUnary(c greetpb.GreetServiceClient) error {
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: dFGreeting,
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "error while calling Greet rpc")
	}
	fmt.Printf("Response from Greet: %v\n", res)
	return nil
}

func doServerStreaming(c greetpb.GreetServiceClient) error {
	fmt.Println("Starting a Server Streaming RPC")

	req := &greetpb.GreetManyTimesRequest{Greeting: dFGreeting}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "error while calling GreetManyTimes RPC")
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Stream has ended
			break
		}
		if err != nil {
			return errors.Wrap(err, "error while reading stream")
		}

		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}

	return nil
}

func doClientStreaming(c greetpb.GreetServiceClient) error {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Elie",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ahuva",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Chanan",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duv",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		return errors.Wrap(err, "error while calling LongGreet")
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(10 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "error while receiving LongGreet response")
	}
	fmt.Printf("LongGreetResponse: %v\n", response)
	return nil
}

func doBiDiStreaming(c greetpb.GreetServiceClient) error {
	fmt.Println("Starting to do BiDi Streaming RPC")

	// Create stream

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to create stream")
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Elie",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ahuva",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Chanan",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Duv",
			},
		},
	}

	// Send messages to client (go routine)
	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(10 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// Receive a bunch of messages from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error!! %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()
	// Block until things are done
	<-waitc

	return nil
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) error {
	fmt.Println("Starting to do Unary GreetWithDeadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: dFGreeting,
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				return errors.Wrap(err, "Timeout was hit! Deadline was exceeded")
			} else {
				return errors.Wrap(err, "unknown error encountered")
			}
		}

		return errors.Wrap(err, "error while calling GreetWithDeadline rpc")
	}
	fmt.Printf("Response from GreetWithDeadline: %v\n", res)
	return nil
}

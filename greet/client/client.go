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
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithInsecure())
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
	return doClientStreaming(c)
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
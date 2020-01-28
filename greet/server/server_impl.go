package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/dfreilich/grpc-samples/greet/greetpb"
	"github.com/pkg/errors"
)

// Server Implementation for Greet RPC Calls
type Server struct{}

// Greet is a unary RPC call to get a name, and send the appropriate greeting
func (s Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	first := req.GetGreeting().GetFirstName()
	last := req.GetGreeting().GetLastName()
	result := "Hello, " + first + " " + last
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

// GreetManyTimes is a Server streaming RPC call
func (s Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s %s! This is your %d message", firstName, lastName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to send response %v", res))
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

// LongGreet is a client-streaming RPC call
func (s Server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Starting LongGreet method")

	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Response: result,
			})
		} else if err != nil {
			return errors.Wrap(err, "error streaming")
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello" + firstName + "!\n "
	}
}

// GreetEveryone is a bi-directional RPC call
func (s Server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("Starting GreetEveryone method")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return errors.Wrap(err, "error while reading the client stream")
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "!"

		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			return errors.Wrap(err, "failed to send data to client")
		}
	}
}

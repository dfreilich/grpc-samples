package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dfreilich/grpc-samples/greet/greetpb"

	"google.golang.org/grpc"

	"github.com/pkg/errors"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	first := req.GetGreeting().GetFirstName()
	last := req.GetGreeting().GetLastName()
	result := "Hello, " + first + " " + last
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}
func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
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

const address = "0.0.0.0"
const port = "50051"

func main() {
	fmt.Println("Running server!")
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
	log.Println("Succesfully ran")
}

func run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

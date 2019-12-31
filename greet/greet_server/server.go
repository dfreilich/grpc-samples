package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dfreilich/grpc-samples/greet/greetpb"

	"google.golang.org/grpc"

	"github.com/pkg/errors"
)

type server struct{}

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

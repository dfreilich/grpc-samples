package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = "0.0.0.0"
const port = "50051"

func main() {
	log.Println("Running calculator RPC server")
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v\n", err)
	}
	log.Println("Succesfully ran.")
}

func run() error {
	listenAddress := fmt.Sprintf("%s:%s", address, port)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to listen on address %s", listenAddress))
	}
	defer lis.Close()

	s := grpc.NewServer()
	defer s.Stop()

	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}
	return nil
}

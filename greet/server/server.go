package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dfreilich/grpc-samples/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/pkg/errors"
)

const address = "0.0.0.0"
const port = "50051"

func main() {
	fmt.Println("Running greet RPC server!")
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
	defer lis.Close()

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return errors.Wrap(err, "failed loading certificates")
	}
	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	defer s.Stop()

	greetpb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

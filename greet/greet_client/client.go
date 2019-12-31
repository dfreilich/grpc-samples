package main

import (
	"fmt"
	"log"

	"github.com/dfreilich/grpc-samples/greet/greetpb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const address = "localhost"
const port = "50051"

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
	fmt.Printf("Created client: %f\n", c)
	return nil
}

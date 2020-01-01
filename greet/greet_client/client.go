package main

import (
	"context"
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
	fmt.Println("Created client")

	return doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) error {
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "David",
			LastName:  "Freilich",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "error while calling Greet rpc")
	}
	fmt.Printf("Response from Greet: %v\n", res)
	return nil
}

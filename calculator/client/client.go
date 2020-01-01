package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const address = "0.0.0.0"
const port = "50051"

func main() {
	log.Println("Running calculator client")
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v\n", err)
	}
	log.Println("Succesfully ran.")
}

func run() error {
	address := fmt.Sprintf("%s:%s", address, port)
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "could not connect to "+address)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	fmt.Println("Created client")

	return doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) error {
	fmt.Println("Starting to do Unary RPC...")
	req := &calculatorpb.SumRequest{
		Nums: []int32{10, 3, 25},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "error while calling Sum RPC")
	}
	fmt.Printf("Response from Sum: %v\n", res.GetSumResult())
	return nil
}

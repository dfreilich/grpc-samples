package main

import (
	"context"
	"fmt"
	"io"
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

	// return doUnarySum(c)
	return doPrimeNumberDecomposition(c)
}

func doUnarySum(c calculatorpb.CalculatorServiceClient) error {
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

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient) error {
	fmt.Println("Starting to do Server Streaming RPC call for Prime Number Decomposition...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Num: int32(1241252343),
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "failed to request prime number decomposition")
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.Wrap(err, "error receiving prime number decomposition")
		}
		fmt.Printf("Next factor: %v\n", res.GetPrimeFactor())
	}

	return nil
}

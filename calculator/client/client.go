package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// return doPrimeNumberDecomposition(c)
	// return doClientStreamingComputeAverage(c)
	// return doBiDiStreamingFindMaximum(c)

	return doErrorUnary(c)
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

func doServerStreamingPrimeNumberDecomp(c calculatorpb.CalculatorServiceClient) error {
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

func doClientStreamingComputeAverage(c calculatorpb.CalculatorServiceClient) error {
	fmt.Println("Starting to do ComputerAverage RPC")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		return errors.Wrap(err, "error creating ComputeAverage stream")
	}

	nums := []int32{3, 5, 9, 54, 23}
	for _, num := range nums {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Num: num,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "failed to receive response for ComputeAverage")
	}
	fmt.Printf("Computed Averaged: %v\n", res.GetAverage())

	return nil
}

func doBiDiStreamingFindMaximum(c calculatorpb.CalculatorServiceClient) error {
	fmt.Println("Starting to do BiDi Streaming for Finding Maximum")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to start Find Maximum stream")
	}

	waitc := make(chan struct{})

	nums := []int32{1, 5, 3, 6, 2, 20}
	go func() {
		for _, num := range nums {
			fmt.Printf("Sending %d\n", num)
			err := stream.Send(&calculatorpb.FindMaximumRequest{
				Num: num,
			})
			if err != nil {
				log.Fatalf("Err! %v", err)
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Failed to receive: %v", err)
				break
			}
			fmt.Printf("Current Max: %d\n", res.GetCurrentMaximum())
		}
		close(waitc)
	}()

	<-waitc
	return nil

}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) error {
	fmt.Println("Starting to do a SquareRoot Unary RPC call")

	// correct call
	num := int32(10)
	err := handleSquareRootRPC(c, num)
	if err != nil {
		return err
	}
	// error call

	num = int32(-2)
	return handleSquareRootRPC(c, num)
}

func handleSquareRootRPC(c calculatorpb.CalculatorServiceClient, num int32) error {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: num})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				return fmt.Errorf("we sent a negative number")
			}
		} else {
			return errors.Wrap(err, "error getting square root with code")
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", num, res.GetNumberRoot())

	return nil
}

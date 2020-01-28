package main

import (
	"context"
	"fmt"
	"io"
	"math"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
)

// Server Implementation for Calculator Proto
type Server struct{}

// Sum is a Unary proto call that adds two numbers
func (s Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	nums := req.GetNums()
	var sum int32
	for _, num := range nums {
		sum += num
	}

	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

// PrimeNumberDecomposition does server side streaming, accepting a request and returning a stream of numbers decomposing the number given to its roots
func (s Server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition was invoked with %v\n", req)
	num := req.GetNum()
	divisor := int32(2)

	for num > 1 {
		if num%divisor == 0 {
			fmt.Printf("Found factor: %v\n", divisor)
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			num = num / divisor
		} else {
			divisor++
		}
	}
	return nil
}

// ComputeAverage is a client-streaming RPC call, accepting numbers and computing the cumulative average of them
func (s Server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Starting ComputeAverage RPC")

	sum := int32(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: avg,
			})
		} else if err != nil {
			return errors.Wrap(err, "error receiving message in ComputeAverage")
		}

		sum += req.GetNum()
		count++
	}
}

// FindMaximum is a bi-directional RPC call, accepting a stream of numbers and returning a number if it is the current maximum
func (s Server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("Starting FindMaximum method")

	maximum := int32(math.MinInt32)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return errors.Wrap(err, "error receiving message")
		}
		num := req.GetNum()
		if num > maximum {
			maximum = num
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				CurrentMaximum: maximum,
			})
			if sendErr != nil {
				return errors.Wrap(err, "failed to send current maximum")
			}
		}
	}
}

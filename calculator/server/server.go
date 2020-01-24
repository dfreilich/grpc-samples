package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const address = "0.0.0.0"
const port = "50051"

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
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

func (s *server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
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

func (s server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
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

func (s server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
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
		maximum = max(maximum, num)

		sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
			CurrentMaximum: maximum,
		})
		if sendErr != nil {
			return errors.Wrap(err, "failed to send current maximum")
		}
	}
}

func max(x, y int32) int32 {
	if x >= y {
		return x
	}
	return y
}

func main() {
	log.Println("Running calculator server")
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

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}
	return nil
}

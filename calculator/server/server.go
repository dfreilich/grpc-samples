package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dfreilich/grpc-samples/calculator/calculatorpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const address = "0.0.0.0"
const port = "50051"

/*

The function takes a Request message that has two integers, and returns a Response that represents the sum of them.
Remember to first implement the service definition in a .proto file, alongside the RPC messages
Implement the Server code first
Test the server code by implementing the Client
Example:

The client will send two numbers (3 and 10) and the server will respond with (13)
*/

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

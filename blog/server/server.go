package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/dfreilich/grpc-samples/blog/blogpb"
	blog "github.com/dfreilich/grpc-samples/blog/server/impl"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = "0.0.0.0"
const port = "50051"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Running Blog RPC server!")
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
	defer fmt.Println("Closing the listener")
	defer lis.Close()

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	blogpb.RegisterBlogServiceServer(s, &blog.Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()

	return nil
}

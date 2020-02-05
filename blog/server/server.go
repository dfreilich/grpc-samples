package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/dfreilich/grpc-samples/blog/blogpb"
	blog "github.com/dfreilich/grpc-samples/blog/server/impl"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const address = "0.0.0.0"
const port = "50051"

var collection *mongo.Collection

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Running Blog RPC server!")
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
	log.Println("Succesfully ran")
}

func run() error {
	fmt.Println("Connecting to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return errors.Wrap(err, "failed to create Mongo client")
	}
	defer client.Disconnect(context.TODO())

	err = client.Connect(context.TODO())
	if err != nil {
		return errors.Wrap(err, "failed to connect to the client")
	}

	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	defer fmt.Println("Closing the listener")
	defer lis.Close()

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	defer s.Stop()
	blogpb.RegisterBlogServiceServer(s, &blog.Server{
		Collection: collection,
	})
	reflection.Register(s)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ch

	return nil
}

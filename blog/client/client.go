package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dfreilich/grpc-samples/blog/blogpb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const address = "localhost"
const port = "50051"

func main() {
	log.Println("Starting blog client")
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

	c := blogpb.NewBlogServiceClient(cc)
	fmt.Println("Created client")

	return doUnary(c)
}

func doUnary(c blogpb.BlogServiceClient) error {
	fmt.Println("Starting to do CreateBlog RPC...")
	in := &blogpb.Blog{
		AuthorId: "David",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	req := &blogpb.CreateBlogRequest{
		Blog: in,
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "error while calling CreateBlog RPC")
	}

	fmt.Printf("Response from CreateBlog: %v\n", res.GetBlog())
	return nil
}

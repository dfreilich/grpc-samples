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

	blog, err := CreateBlog(c)
	if err != nil {
		return errors.Wrap(err, "failed to create blog")
	}

	if err := ReadBlog(c, blog); err != nil {
		return errors.Wrap(err, "failed to read blog")
	}

	blog.Title = "My First Blog (edited)"
	if err := UpdateBlog(c, blog); err != nil {
		return errors.Wrap(err, "failed to update blog")
	}

	if err := DeleteBlog(c, blog.GetId()); err != nil {
		return errors.Wrap(err, "failed to delete blog")
	}

	return nil
}

// CreateBlog creates a blog post
func CreateBlog(c blogpb.BlogServiceClient) (*blogpb.Blog, error) {
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
		return nil, errors.Wrap(err, "error while calling CreateBlog RPC")
	}

	fmt.Printf("Response from CreateBlog: %v\n", res.GetBlog())
	return res.GetBlog(), nil
}

// ReadBlog reads a blog post
func ReadBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog) error {
	req := &blogpb.ReadBlogRequest{BlogId: "fake id"}
	_, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blog.GetId()})
	if err != nil {
		return errors.Wrap(err, "failed to read blog")
	}

	fmt.Printf("Blog was read: %v\n", res)
	return nil
}

// UpdateBlog updates a given blog with the new blog fields
func UpdateBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog) error {
	fmt.Printf("Updating blog with: %v\n", blog)
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		return errors.Wrap(err, "failed to update blog")
	}

	fmt.Printf("Blog was updated: %v\n", res.GetBlog())
	return nil
}

// DeleteBlog updates a given blog with the new blog fields
func DeleteBlog(c blogpb.BlogServiceClient, id string) error {
	fmt.Printf("Deleting blog with id: %v\n", id)
	_, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: id})
	if err != nil {
		return errors.Wrap(err, "failed to update blog")
	}

	fmt.Printf("Blog was deleted\n")
	return nil
}

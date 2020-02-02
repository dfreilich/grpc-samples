package blog

import (
	"context"
	"fmt"

	"github.com/dfreilich/grpc-samples/blog/blogpb"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Item is an itemized blog post
type Item struct {
	ID       primitive.ObjectID `bson:"_d, omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

// Server is an implementation of the Blog RPC proto
type Server struct {
	Collection *mongo.Collection
}

// CreateBlog creates a blog entry
func (s Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()

	data := Item{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := s.Collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to insert %v into collection with err: %v", data, err))
	}

	objectID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Can't convert to OID")
	}

	return &blogpb.CreateBlogResponse{
		Blog: convertItemToBlog(data, objectID.Hex()),
	}, nil
}

// ReadBlog takes in a blog id, and returns a Blog
func (s Server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("ReadBlog Request Received")

	id := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", id),
		)
	}

	blog := &Item{}
	filter := bson.M{"_id": oid}
	res := s.Collection.FindOne(context.Background(), filter)
	if err := res.Decode(blog); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", id),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: convertItemToBlog(*blog, ""),
	}, nil
}

// UpdateBlog updates a blog post
func (s Server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Update blog request")
	blog := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", blog.GetId()),
		)
	}

	data := &Item{}
	filter := bson.M{"_id": oid}
	res := s.Collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", oid),
		)
	}

	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()
	_, err = s.Collection.ReplaceOne(context.Background(), filter, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", err),
		)
	}
	fmt.Printf("Blog was updated with response: %v\n", data)

	return &blogpb.UpdateBlogResponse{
		Blog: convertItemToBlog(*data, ""),
	}, nil
}

// DeleteBlog deletes a blog by blog_id
func (s Server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blog request")
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", req.GetBlogId()),
		)
	}

	filter := bson.M{"_id": oid}
	res, err := s.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", oid),
		)
	}

	return &blogpb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
}

func convertItemToBlog(item Item, id string) *blogpb.Blog {
	blog := blogpb.Blog{
		Id:       item.ID.Hex(),
		AuthorId: item.AuthorID,
		Title:    item.Title,
		Content:  item.Content,
	}

	if id != "" {
		blog.Id = id
	}

	return &blog
}

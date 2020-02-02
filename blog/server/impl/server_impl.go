package blog

import (
	"context"
	"fmt"

	"github.com/dfreilich/grpc-samples/blog/blogpb"
	"github.com/mongodb/mongo-go-driver/mongo"
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
		Blog: &blogpb.Blog{
			Id	:       objectID.Hex(),
			AuthorId: blog.GetAuthorId(),
			Content:  blog.GetContent(),
			Title:    blog.GetTitle(),
		},
	}, nil
}

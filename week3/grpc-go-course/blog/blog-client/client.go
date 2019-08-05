package main 

import (
	"log"
	"io"
	"context"
	"google.golang.org/grpc"
	"github.com/nhaancs/grabvn-golang-bootcamp/week3/grpc-go-course/blog/blogpb"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	// createBlog(c)
	// readBlog(c)
	// updateBlog(c)
	// deleteBlog(c)
	listBlog(c)
}

func createBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "qwerty1",
			Title: "Title 1",
			Content: "Content 1",
		},
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	log.Println("created blog: ", res.GetBlog())
}

func readBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.ReadBlogRequest{ Id: "5d47f49762c4f3e2284dec18" }
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	log.Println("read blog: ", res.GetBlog())
}

func updateBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id: "5d47f49762c4f3e2284dec18",
			AuthorId: "qwerty1 edited",
			Title: "Title 1 edited",
			Content: "Content 1 edited",
		},
	}
	res, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	log.Println("updated blog: ", res.GetBlog())
}

func deleteBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.DeleteBlogRequest{ Id: "5d47f49762c4f3e2284dec18" }
	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	log.Println("deleted blog: ", res.GetId())
}

func listBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.ListBlogRequest{}

	resStream, err := c.ListBlog(context.Background(), req)
	if err != nil { 
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}

	log.Println("list blogs: ")
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v", err)
		}
		log.Println(msg.GetBlog())
	}
}
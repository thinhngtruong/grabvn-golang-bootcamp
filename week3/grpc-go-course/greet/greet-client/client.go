package main 

import (
	"log"
	"fmt"
	"context"
	"io"
	"time"
	"google.golang.org/grpc"
	"github.com/nhaancs/grabvn-golang-bootcamp/week3/grpc-go-course/greet/greetpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// doUnary(c)	
	// doServeStreaming(c)
	// doClientStreaming(c)
	doBidirectionalStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nhan",
			LastName: "Nguyen",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeting: %v", err)
	}

	fmt.Println(res.Result)
}

func doServeStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nhan",
			LastName: "Nguyen",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil { 
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v", err)
		}
		log.Println(msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Nhan",
				LastName: "Nguyen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Phuong",
				LastName: "Nguyen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hieu",
				LastName: "Nguyen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Le",
				LastName: "Nguyen",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tri",
				LastName: "Nguyen",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil { 
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil { 
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}
	log.Println(res)
}

func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Nhan",
				LastName: "Nguyen",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Phuong",
				LastName: "Nguyen",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hieu",
				LastName: "Nguyen",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Le",
				LastName: "Nguyen",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tri",
				LastName: "Nguyen",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("failed while calling GreetEveryone RPC: %v", err)
		return
	}

	wait := make(chan bool)

	go func() {
		for _, req := range requests {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("failed while reading stream: %v", err)
				break
			}
			log.Println(res.GetResult())
		}

		close(wait)
	}()

	<-wait
}
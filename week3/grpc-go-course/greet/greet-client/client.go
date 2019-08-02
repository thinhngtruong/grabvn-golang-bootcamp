package main 

import (
	"log"
	"fmt"
	"context"
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
	// fmt.Printf("created client: %v\n", c)

	doUnary(c)	
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
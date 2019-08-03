package main

import (
	"log"
	"fmt"
	"context"
	"io"
	"google.golang.org/grpc"
	"github.com/nhaancs/grabvn-golang-bootcamp/week3/grpc-go-course/calculator/calculatorpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doCalculate(c)
	doPrimeNumberDecomposition(c)
}

func doCalculate(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Calculating: &calculatorpb.Calculating{
			Num1: 10,
			Num2: 20,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum: %v", err)
	}

	fmt.Println(res.Result)
}

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 105,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil { 
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v", err)
		}
		log.Println(msg.GetResult())
	}
}
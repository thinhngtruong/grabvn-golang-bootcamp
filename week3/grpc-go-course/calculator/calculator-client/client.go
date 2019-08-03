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
	// doPrimeNumberDecomposition(c)
	doComputeAverage(c)
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

func doComputeAverage(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 10,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 11,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 12,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 13,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 14,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil { 
		log.Fatalf("error while calling ComputeAverage RPC: %v", err)
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
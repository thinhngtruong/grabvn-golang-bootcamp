package main

import (
	"log"
	"fmt"
	"context"
	"io"
	"time"
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
	// doComputeAverage(c)
	doFindMaximum(c)
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

func doFindMaximum(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.FindMaximumRequest {
		&calculatorpb.FindMaximumRequest{
			Number: 1,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 2,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 3,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 4,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 5,
		},
	}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("failed while calling GreetEveryone RPC: %v", err)
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
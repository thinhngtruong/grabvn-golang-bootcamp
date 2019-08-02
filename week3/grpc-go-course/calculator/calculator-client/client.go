package main

import (
	"log"
	"fmt"
	"context"
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
	doCalculate(c)
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
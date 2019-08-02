package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/nhaancs/grabvn-golang-bootcamp/week3/grpc-go-course/calculator/calculatorpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	num1 := req.GetCalculating().GetNum1()
	num2 := req.GetCalculating().GetNum2()

	res := &calculatorpb.SumResponse{
		Result: num1 + num2,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to  listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
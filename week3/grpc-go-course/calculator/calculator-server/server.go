package main

import (
	"context"
	"log"
	"net"
	"math"
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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	number := req.GetNumber() 
	sendResponse := func(result int32) {
		res := &calculatorpb.PrimeNumberDecompositionResponse{
			Result: result,
		}
		stream.Send(res)
	}

	// Print the number of two's that divide n 
	for number % 2 == 0 {
		sendResponse(2)
		number /= 2
	}

	// n must be odd at this point 
	// so a skip of 2 ( i = i + 2) can be used 
	var i int32 = 3
	for ; float64(i) <= math.Sqrt(float64(number)); i += 2 {
		for number % i == 0 {
			sendResponse(i)
			number /= i
		}
	}

	// Condition if n is a prime 
	// number greater than 2 
	if number > 2 {
		sendResponse(number)
	}

	return nil
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
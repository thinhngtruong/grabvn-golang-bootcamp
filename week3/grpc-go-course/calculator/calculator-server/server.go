package main

import (
	"context"
	"log"
	"net"
	"math"
	"io"
	"fmt"
	"google.golang.org/grpc"
	"github.com/nhaancs/grabvn-golang-bootcamp/week3/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	var numbers []int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var total int32
			for _, v := range numbers {
				total += v
			}
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: float32(total/int32(len(numbers))),
			})
		}

		if err != nil {
			log.Fatalf("failed while reading stream: %v", err)
		}

		number := req.GetNumber()
		numbers = append(numbers, number)
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	maximum := int32(0)
	for {
		req, err := stream.Recv() 
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed while reading stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number

			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Result: number,
			})
			if err != nil {
				log.Fatalf("failed while sending to client: %v", err)
				return err
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received a nagetive number: %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to  listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	
	// register reflection service on grpc server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
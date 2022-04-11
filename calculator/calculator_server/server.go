package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tobyartisan-others-tutorials/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPDC: %v\n", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func (*server) DecomposeToPrimeNumbers(req *calculatorpb.DecomposeToPrimeNumbersRequest, stream calculatorpb.CalculatorService_DecomposeToPrimeNumbersServer) error {
	fmt.Printf("Received DecomposeToPrimeNumbers RPC: %v\n", req)
	num := req.GetNumber()

	divisor := int64(2)
	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&calculatorpb.DecomposeToPrimeNumbersResponse{
				PrimeFactor: divisor,
			})
			num = num / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}

	return nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

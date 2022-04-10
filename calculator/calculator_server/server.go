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
	num := req.GetNumber()
	var primes []int32

	k := int32(2)
	for num > 1 {
		if num%k == 0 {
			primes = append(primes, k)
			num = num / k

			res := &calculatorpb.DecomposeToPrimeNumbersResponse{
				PrimeNumber: k,
			}
			stream.Send(res)
		} else {
			k = k + 1
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

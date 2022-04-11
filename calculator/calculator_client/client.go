package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/tobyartisan-others-tutorials/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//doUnary(c)

	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 4,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Prime Decomposition Server Streaming RPC...")
	req := &calculatorpb.DecomposeToPrimeNumbersRequest{
		Number: 12390392840,
	}
	stream, err := c.DecomposeToPrimeNumbers(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Prime Decompositionn RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)

		}
		fmt.Println(res.GetPrimeFactor())
	}
}

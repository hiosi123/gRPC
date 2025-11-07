package main

import (
	"context"
	"log"

	pb "github.com/hiosi123/gRPC/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("do sum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 1,
		SecondNumer: 2,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("result of Sum: %v\n", res.Result)
}

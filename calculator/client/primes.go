package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hiosi123/gRPC/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Printf("doPrime initinated")

	stream, err := c.GetPrime(context.Background(), &pb.PrimeRequest{
		PrimeNumber: int32(123123122),
	})

	if err != nil {
		log.Fatalf("error while calliing getPrime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("getPrime: %v\n", msg.Result)
	}
}

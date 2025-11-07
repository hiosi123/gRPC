package main

import (
	"context"
	"log"

	pb "github.com/hiosi123/gRPC/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

func doSum(c pb.GreetServiceClient) {
	log.Println("do sum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstValue:  1,
		SecondValue: 2,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("result of Sum: %v\n", res.SumResult)
}

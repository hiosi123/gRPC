package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hiosi123/gRPC/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Simon"},
		{FirstName: "Hanna"},
		{FirstName: "Chist"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreeL %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sendiong req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hiosi123/gRPC/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg initiated")

	reqs := []*pb.AvgRequest{
		{Num: 1},
		{Num: 100},
		{Num: 50},
		{Num: 40},
		{Num: 30},
	}

	stream, err := c.GetAvg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receving response from avg %v\n", err)
	}
	log.Printf("avg: %v\n", res.Result)
}

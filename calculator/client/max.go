package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/hiosi123/gRPC/calculator/proto"
)

func doCurrentMax(c pb.CalculatorServiceClient) {
	log.Println("doCurrentMax was invoked")

	stream, err := c.GetCurrentMax(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Num: 1},
		{Num: 5},
		{Num: 4},
		{Num: 3},
		{Num: 2},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receving: %v\n", err)
				break
			}
			log.Printf("received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}

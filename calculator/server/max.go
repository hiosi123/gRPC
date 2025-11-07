package main

import (
	"io"
	"log"
	"math"

	pb "github.com/hiosi123/gRPC/calculator/proto"

	"google.golang.org/grpc"
)

func (s *Server) GetCurrentMax(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Printf("GerCurrentMax initiated")

	currentMax := math.MinInt64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading clients stream: %v\n", err)
		}

		if currentMax < int(req.Num) {
			currentMax = int(req.Num)
		}

		err = stream.Send(&pb.MaxResponse{
			Result: int64(currentMax),
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}

}

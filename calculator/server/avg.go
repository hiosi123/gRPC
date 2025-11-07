package main

import (
	"io"
	"log"

	pb "github.com/hiosi123/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) GetAvg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Printf("GetAvg server streaming invoked")

	count := 0
	sum := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			result := sum / count

			return stream.SendAndClose(&pb.AvgResponse{
				Result: int64(result),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sum += int(res.Num)
		count++
	}
}

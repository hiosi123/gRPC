package main

import (
	"log"

	pb "github.com/hiosi123/gRPC/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) GetPrime(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("GetPrime function invoked from server with %v\n", in)

	if in.PrimeNumber == 0 {
		log.Printf("wrong number")
		stream.Context().Err()
	}

	calculatePrimeDivision(int(in.PrimeNumber), stream)

	return nil
}

// 120,
// 2,2,2,3,5
func calculatePrimeDivision(num int, stream grpc.ServerStreamingServer[pb.PrimeResponse]) {
	k := 2
	N := num

	for num > 1 {
		if N%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: int32(k),
			})
			N = N / k
		} else {
			k = k + 1
		}
	}
}

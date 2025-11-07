package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hiosi123/gRPC/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrime(c)
	// doAvg(c)
	// doCurrentMax(c)
	doSqrt(c, -2)

}

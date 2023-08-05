package main

import (
	"log"
	"net"
	pb "github.com/tanayvaswani/grpc-impl-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
} 

func main() {
	lisn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	// gRPC Server 
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lisn); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
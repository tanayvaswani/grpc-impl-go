package main

import (
	"log"
	pb "github.com/tanayvaswani/grpc-impl-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection to the server failed: %v", err)
	}
	// closing the client
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Tanay1", "Tanay2", "Tanay3"},
	// }

	// callSayHello(client)
}
package main

import (
	"log"

	pb "github.com/MahadevMaahi/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost" + PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to the server %v", err)
	}
	defer conn.Close();

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList {
		Names: []string{"Alice", "Bob", "carl"},
	}

	callSayHello(client)
	CallSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
}
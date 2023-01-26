package main

import (
	"context"
	"log"
	"time"
	pb "github.com/MahadevMaahi/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not Start Client Stream %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest {
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error While Sending %v", err)
		}
		log.Printf("Send the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("Error While Receiving %v", err)
	}
	log.Printf("Received Response %v", res.Messages)
}
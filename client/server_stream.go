package main

import (
	"io"
	"context"
	"log"
	pb "github.com/MahadevMaahi/go-grpc/proto"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming has Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names ; %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error Receiving Stream %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
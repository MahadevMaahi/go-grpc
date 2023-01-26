package main

import (
	"log"
	"io"
	pb "github.com/MahadevMaahi/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			log.Fatalf("Error While Receiving the Stream %v", err)
			return err;
		}
		log.Printf("Received a Request with name :%v", req.Name)
		messages = append(messages, "hello from " + req.Name)
	}
}
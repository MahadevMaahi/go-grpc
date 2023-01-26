package main

import (
	"log"
	"context"
	"time"

	pb "github.com/MahadevMaahi/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("could not Greet %v", err)
	}

	log.Printf("%s", res.Message)
}
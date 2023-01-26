package main

import (
	"context"
	pb "github.com/MahadevMaahi/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse {
		Message : "Hello There!",
	}, nil
}
package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/printqueue/proto"
)

func (s *Server) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	return &pb.PrintResponse{}, status.Errorf(codes.Unimplemented, "Not ready for this yet")
}

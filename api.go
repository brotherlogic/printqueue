package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pb "github.com/brotherlogic/printqueue/proto"
	rspb "github.com/brotherlogic/rstore/proto"
	"github.com/google/uuid"
)

func (s *Server) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	stored := &pb.StoredPrintRequest{
		Lines:       req.GetLines(),
		Origin:      req.GetOrigin(),
		Urgency:     req.GetUrgency(),
		Destination: req.GetDestination(),
		Fanout:      req.GetFanout(),
	}
	data, _ := proto.Marshal(stored)

	uid := uuid.New().String()
	_, err := s.client.Write(ctx, &rspb.WriteRequest{
		Key:   fmt.Sprintf("printqueue/%v", uid),
		Value: &anypb.Any{Value: data},
	})

	return &pb.PrintResponse{Id: uid}, err
}

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

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, err := s.client.Delete(ctx, &rspb.DeleteRequest{
		Key: fmt.Sprintf("printqueue/%v", req.GetId()),
	})
	return &pb.DeleteResponse{}, err
}

func (s *Server) RegisterPrinter(ctx context.Context, req *pb.RegisterPrinterRequest) (*pb.RegisterPrinterResponse, error) {
	s.printers = append(s.printers, &printer{
		id:      req.GetId(),
		address: req.GetCallbackAddress(),
		ptype:   req.GetReceiverType(),
	})
	return &pb.RegisterPrinterResponse{}, nil
}

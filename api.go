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
	uid := uuid.New().String()

	stored := &pb.StoredPrintRequest{
		Lines:       req.GetLines(),
		Origin:      req.GetOrigin(),
		Urgency:     req.GetUrgency(),
		Destination: req.GetDestination(),
		Fanout:      req.GetFanout(),
	}
	data, _ := proto.Marshal(stored)

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

func (s *Server) Heartbeat(_ context.Context, _ *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{}, nil
}

func (s *Server) Ack(ctx context.Context, req *pb.AckRequest) (*pb.AckResponse, error) {
	job, err := s.client.Read(ctx, &rspb.ReadRequest{
		Key: fmt.Sprintf("printqueue/%v", req.GetId()),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read entry %v -> %w", req.GetId(), err)
	}

	val := &pb.StoredPrintRequest{}
	err = proto.Unmarshal(job.GetValue().GetValue(), val)
	if err != nil {
		return nil, err
	}

	if val.GetDestination() == req.GetAckType() {
		if val.GetFanout() == pb.Fanout_FANOUT_ONE {
			_, err = s.client.Delete(ctx,
				&rspb.DeleteRequest{
					Key: fmt.Sprintf("printqueue/%v", req.GetId()),
				})
			return &pb.AckResponse{}, err
		}
	}

	return &pb.AckResponse{}, nil
}

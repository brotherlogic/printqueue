package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pb "github.com/brotherlogic/printqueue/proto"
	rspb "github.com/brotherlogic/rstore/proto"
	"github.com/google/uuid"
)

func (s *Server) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	if len(req.GetLines()) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "you need to have something to print: %v", req)
	}

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
	if err != nil {
		return nil, fmt.Errorf("unable to write print request to queue: %w", err)
	}

	return &pb.PrintResponse{Id: uid}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, err := s.client.Delete(ctx, &rspb.DeleteRequest{
		Key: fmt.Sprintf("printqueue/%v", req.GetId()),
	})
	return &pb.DeleteResponse{}, err
}

func convertToPrintJob(elem *pb.StoredPrintRequest, id string) *pb.PrintJob {
	// The id here contains the  directory prefix - so strip that off
	return &pb.PrintJob{
		Lines:   elem.GetLines(),
		Urgency: elem.GetUrgency(),
		Id:      id[len("printqueue/"):],
	}
}

func (s *Server) RegisterPrinter(ctx context.Context, req *pb.RegisterPrinterRequest) (*pb.RegisterPrinterResponse, error) {
	s.printers = append(s.printers, &printer{
		id:      req.GetId(),
		address: req.GetCallbackAddress(),
		ptype:   req.GetReceiverType(),
	})

	queue, err := s.getQueue(ctx)
	if err != nil {
		return nil, err
	}

	var rqueue []*pb.PrintJob
	for id, elem := range queue {
		if elem.GetDestination() == req.GetReceiverType() {
			rqueue = append(rqueue, convertToPrintJob(elem, id))
		}
	}

	return &pb.RegisterPrinterResponse{
		Jobs: rqueue,
	}, nil
}

func (s *Server) Heartbeat(_ context.Context, _ *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{}, nil
}

func (s *Server) Ack(ctx context.Context, req *pb.AckRequest) (*pb.AckResponse, error) {
	if req.GetAckType() == pb.Destination_DESTINATION_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "You must inclue an ack type")
	}

	job, err := s.client.Read(ctx, &rspb.ReadRequest{
		Key: fmt.Sprintf("printqueue/%v", req.GetId()),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to read the entry %v -> %w", req.GetId(), err)
	}

	val := &pb.StoredPrintRequest{}
	err = proto.Unmarshal(job.GetValue().GetValue(), val)
	if err != nil {
		return nil, err
	}

	if val.GetDestination() == req.GetAckType() {
		if val.GetFanout() == pb.Fanout_FANOUT_ONE || val.GetFanout() == pb.Fanout_FANOUT_UNKNOWN {
			_, err = s.client.Delete(ctx,
				&rspb.DeleteRequest{
					Key: fmt.Sprintf("printqueue/%v", req.GetId()),
				})
			return &pb.AckResponse{}, err
		} else {
			log.Printf("Skipping %v", val.GetFanout())
		}
	} else {
		log.Printf("Skipping %v and %v", val.GetDestination(), req.GetAckType())
	}

	return &pb.AckResponse{}, nil
}

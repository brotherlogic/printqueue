package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/printqueue/proto"
	rstore_client "github.com/brotherlogic/rstore/client"
)

func InitTestServer() *Server {
	return &Server{client: rstore_client.GetTestClient()}
}

func TestPrint(t *testing.T) {
	s := InitTestServer()

	res, err := s.Print(context.Background(), &pb.PrintRequest{})
	if err != nil || res.GetId() == "" {
		t.Errorf("Failed to print: %v and %v", res, err)
	}
}

func TestAck(t *testing.T) {
	s := InitTestServer()

	res, err := s.Print(context.Background(), &pb.PrintRequest{
		Fanout:      pb.Fanout_FANOUT_ONE,
		Destination: pb.Destination_DESTINATION_RECEIPT,
	})
	if err != nil || res.GetId() == "" {
		t.Errorf("Failed to print: %v and %v", res, err)
	}

	jobs, err := s.RegisterPrinter(context.Background(), &pb.RegisterPrinterRequest{
		ReceiverType: pb.Destination_DESTINATION_RECEIPT,
	})
	if err != nil || len(jobs.GetJobs()) != 1 {
		t.Fatalf("Bad register: %v and %v", err, jobs)
	}

	_, err = s.Ack(context.Background(), &pb.AckRequest{
		Id: jobs.GetJobs()[0].GetId(),
	})
	if err != nil {
		t.Errorf("Bad ack: %v", err)
	}

	queue, err := s.getQueue(context.Background())
	if err != nil {
		t.Fatalf("Bad queue get: %v", err)
	}

	if len(queue) != 0 {
		t.Errorf("Print job was not removed post ack: %v (%v)", queue, len(queue))
	}
}

func TestDelete(t *testing.T) {
	s := InitTestServer()

	res, err := s.Print(context.Background(), &pb.PrintRequest{})
	if err != nil || res.GetId() == "" {
		t.Errorf("Failed to print: %v and %v", res, err)
	}

	_, err = s.Delete(context.Background(), &pb.DeleteRequest{Id: res.GetId()})
	if err != nil {
		t.Errorf("Delete fail: %v", err)
	}
}

func TestRegister(t *testing.T) {
	s := InitTestServer()

	_, err := s.RegisterPrinter(context.Background(), &pb.RegisterPrinterRequest{
		Id:              "id",
		CallbackAddress: "fake",
		ReceiverType:    pb.Destination_DESTINATION_UNKNOWN,
	})
	if err != nil {
		t.Errorf("Bad register: %v", err)
	}

	if len(s.getPrinters()) != 1 {
		t.Errorf("Printer was not registered")
	}
}

func TestHeartbeat(t *testing.T) {
	s := InitTestServer()
	_, err := s.Heartbeat(context.Background(), &pb.HeartbeatRequest{})

	if err != nil {
		t.Errorf("Heartbeat failed")
	}
}

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

package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/printqueue/proto"
)

func InitTestServer() *Server {
	return &Server{}
}

func TestPrint(t *testing.T) {
	s := InitTestServer()

	res, err := s.Print(context.Background(), &pb.PrintRequest{})
	if err == nil {
		t.Errorf("Should have failed at least a little: %v", res)
	}
}

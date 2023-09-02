package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/printqueue/proto"
)

type PrintQueueClient struct {
	client pb.PrintServiceClient
}

func NewPrintQueueClient(ctx context.Context) (*PrintQueueClient, error) {
	conn, err := grpc.Dial("print.brotherlogic-backend.com:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing printer: %v", err)
	}

	client := pb.NewPrintServiceClient(conn)

	return &PrintQueueClient{client: client}, nil
}

func (p *PrintQueueClient) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	return p.client.Print(ctx, req)
}

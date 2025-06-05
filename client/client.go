package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/printqueue/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	printErrors = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "printqueue_errors",
		Help: "print queue resposnes",
	}, []string{"code"})
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
	val, err := p.client.Print(ctx, req)
	printErrors.With(prometheus.Labels{"code": fmt.Sprintf("%v", status.Code(err))}).Inc()
	return val, err
}

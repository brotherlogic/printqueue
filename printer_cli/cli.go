package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/brotherlogic/printqueue/proto"
)

func main() {
	conn, err := grpc.Dial("print.brotherlogic-backend.com:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error dialing printer: %v", err)
	}

	client := pb.NewPrintServiceClient(conn)

	_, err = client.Print(context.Background(), &pb.PrintRequest{
		Lines:       []string{"Hello"},
		Destination: pb.Destination_DESTINATION_RECEIPT,
	})
	if err != nil {
		log.Fatalf("bad print: %v", err)
	}

}

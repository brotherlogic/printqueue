package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/printqueue/proto"
)

func main() {
	conn, err := grpc.Dial("printer.brotherlogic-backend.com:80")
	if err != nil {
		log.Fatalf("Error dialing printer: %v", err)
	}

	client := pb.NewPrintServiceClient(conn)

	_, err = client.Print(context.Background(), &pb.PrintRequest{
		Lines:       []string{"Hello"},
		Destination: pb.Destination_DESTINATION_RECEIPT,
	})
	if err != nil {
		log.Fatalf("Bad print: %v", err)
	}

}

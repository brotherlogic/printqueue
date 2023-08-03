package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/printqueue/proto"

	rstore_client "github.com/brotherlogic/rstore/client"
)

var (
	port        = flag.Int("port", 8080, "The server port for grpc traffic")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
)

type Server struct {
	client rstore_client.RStoreClient
}

func NewServer() *Server {
	return &Server{}
}

func main() {
	flag.Parse()

	s := NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("printqueue is unable to listen on the grpc port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterPrintServiceServer(gs, s)

	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("printqueue is unable to serve grpc: %v", err)
		}
		log.Fatalf("printqueue has closed the grpc port")
	}()

	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	log.Fatalf("printqueue is unable to serve metrics: %v", err)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "github.com/brotherlogic/printqueue/proto"
	rspb "github.com/brotherlogic/rstore/proto"

	rstore_client "github.com/brotherlogic/rstore/client"
)

var (
	port        = flag.Int("port", 8080, "The server port for grpc traffic")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")

	queueLen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "printqueue_qlen",
		Help: "The length of the working queue I think yes",
	})
)

type printer struct {
	id      string
	address string
	//lastHeartbeat time.Time
	ptype pb.Destination
}

type Server struct {
	client   rstore_client.RStoreClient
	printers []*printer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) getPrinters() []*printer {
	return s.printers
}

func (s *Server) getQueue(ctx context.Context) ([]*pb.StoredPrintRequest, error) {
	keys, err := s.client.GetKeys(ctx, &rspb.GetKeysRequest{Prefix: "printqueue/"})
	if err != nil {
		return nil, fmt.Errorf("unable to read keys: %w", err)
	}

	queueLen.Set(float64(len(keys.GetKeys())))

	var stored []*pb.StoredPrintRequest
	for _, key := range keys.GetKeys() {
		data, err := s.client.Read(ctx, &rspb.ReadRequest{Key: key})
		if err != nil {
			return nil, fmt.Errorf("unable to read key %v -> %w", key, err)
		}

		val := &pb.StoredPrintRequest{}
		err = proto.Unmarshal(data.GetValue().GetValue(), val)
		if err != nil {
			return nil, fmt.Errorf("cannot read stored data: %v -> %w", key, err)
		}
		stored = append(stored, val)
	}

	return stored, nil
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
			log.Fatalf("printqueue is actually unable to listen on the grpc port: %v", err)
		}
		log.Fatalf("printqueue has closed the grpc port")
	}()

	_, err = s.getQueue(context.Background())
	if err != nil {
		log.Printf("Error getting queue: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	log.Fatalf("printqueue is unable to serve metrics: %v", err)
}

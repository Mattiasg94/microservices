package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Mattiasg94/microservices-proto/golang/seatallocation"
	"github.com/Mattiasg94/microservices/seatallocation/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	seatallocation.UnimplementedSeatAllocationServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		slog.Error("failed to listen on port", "port", a.port, "error", err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	seatallocation.RegisterSeatAllocationServiceServer(grpcServer, a)
	reflection.Register(grpcServer)

	slog.Info("starting seat allocation service on port", "port", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("failed to serve grpc", "port", a.port)
	}
}

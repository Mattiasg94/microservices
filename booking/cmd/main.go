package main

import (
	"log/slog"

	"github.com/Mattiasg94/microservices/booking/internal/adapters/grpc"
	"github.com/Mattiasg94/microservices/booking/internal/adapters/repository"
	"github.com/Mattiasg94/microservices/booking/internal/adapters/seatallocation"
	"github.com/Mattiasg94/microservices/booking/internal/application/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dsn := "file:memdb1?mode=memory&cache=shared"

	repoAdapter, err := repository.NewAdapter(dsn)
	if err != nil {
		slog.Error("failed to initialize repository adapter", "error", err)
	}

	seatAllocationAdapter, err := seatallocation.NewAdapter("localhost:50052")
	if err != nil {
		slog.Error("failed to initialize seat allocation adapter", "error", err)
	}
	application := api.NewApplication(repoAdapter, seatAllocationAdapter)

	grpcAdapter := grpc.NewAdapter(application, 50051)
	grpcAdapter.Run()
}

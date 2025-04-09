package main

import (
	"log"

	"github.com/Mattiasg94/microservices/seatallocation/internal/adapters/grpc"
	"github.com/Mattiasg94/microservices/seatallocation/internal/adapters/repository"
	"github.com/Mattiasg94/microservices/seatallocation/internal/application/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dsn := "file:memdb1?mode=memory&cache=shared"

	repoAdapter, err := repository.NewAdapter(dsn)
	if err != nil {
		log.Fatalf("failed to initialize repository adapter: %v", err)
	}
	seatService := api.NewApplication(repoAdapter)

	grpcServer := grpc.NewAdapter(seatService, 50052)
	grpcServer.Run()
}

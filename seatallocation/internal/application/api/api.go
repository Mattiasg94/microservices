package api

import (
	"context"

	"github.com/Mattiasg94/microservices/seatallocation/internal/application/domain"
	"github.com/Mattiasg94/microservices/seatallocation/internal/ports"
)

type Application struct {
	repo ports.SeatAllocationRepositoryPort
}

func NewApplication(repo ports.SeatAllocationRepositoryPort) *Application {
	return &Application{repo: repo}
}

func (a *Application) AllocateSeat(ctx context.Context, seatAllocation domain.SeatAllocation) error {
	err := a.repo.Save(ctx, seatAllocation)
	if err != nil {
		return err
	}

	return nil
}

package ports

import (
	"context"

	"github.com/Mattiasg94/microservices/seatallocation/internal/application/domain"
)

type SeatAllocationRepositoryPort interface {
	Save(ctx context.Context, seatAllocation domain.SeatAllocation) error
}

package ports

import (
	"context"

	"github.com/Mattiasg94/microservices/seatallocation/internal/application/domain"
)

type APIPort interface {
	AllocateSeat(ctx context.Context, seatAllocation domain.SeatAllocation) error
}

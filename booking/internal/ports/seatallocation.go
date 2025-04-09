package ports

import (
	"context"

	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
)

type SeatAllocationPort interface {
	AllocateSeat(ctx context.Context, booking *domain.Booking) error
}

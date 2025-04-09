package ports

import (
	"context"

	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
)

type BookingRepositoryPort interface {
	Create(ctx context.Context, booking domain.Booking) (domain.Booking, error)
}

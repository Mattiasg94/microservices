package api

import (
	"context"

	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
	"github.com/Mattiasg94/microservices/booking/internal/ports"
)

type Application struct {
	repo           ports.BookingRepositoryPort
	seatAllocation ports.SeatAllocationPort
}

func NewApplication(repo ports.BookingRepositoryPort, seatAllocation ports.SeatAllocationPort) *Application {
	return &Application{repo: repo, seatAllocation: seatAllocation}
}

func (a *Application) CreateBooking(ctx context.Context, booking domain.Booking) (domain.Booking, error) {
	booking, err := a.repo.Create(ctx, booking)
	if err != nil {
		return domain.Booking{}, err
	}

	err = a.seatAllocation.AllocateSeat(ctx, &booking)
	if err != nil {
		return domain.Booking{}, err
	}
	return booking, nil
}

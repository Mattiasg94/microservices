package grpc

import (
	"context"

	"github.com/Mattiasg94/microservices-proto/golang/booking"
	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
)

func (a Adapter) CreateBooking(ctx context.Context, req *booking.CreateBookingRequest) (*booking.CreateBookingResponse, error) {
	createdBooking, _ := a.api.CreateBooking(ctx, domain.Booking{
		CustomerID: req.CustomerId,
		FlightID:   req.FlightId,
		SeatNumber: req.SeatNumber,
	})

	return &booking.CreateBookingResponse{
		BookingId: createdBooking.ID,
	}, nil
}

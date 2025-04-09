package seatallocation

import (
	"context"

	"github.com/Mattiasg94/microservices-proto/golang/seatallocation"
	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	seatAllocationClient seatallocation.SeatAllocationServiceClient
}

func NewAdapter(seatAllocationServiceAddress string) (*Adapter, error) {
	conn, err := grpc.NewClient(seatAllocationServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := seatallocation.NewSeatAllocationServiceClient(conn)
	return &Adapter{seatAllocationClient: client}, nil
}

func (a *Adapter) AllocateSeat(ctx context.Context, booking *domain.Booking) error {
	_, err := a.seatAllocationClient.AllocateSeat(ctx, &seatallocation.AllocateSeatRequest{
		FlightId:   booking.FlightID,
		SeatNumber: booking.SeatNumber,
	})
	if err != nil {
		return err
	}
	return nil
}

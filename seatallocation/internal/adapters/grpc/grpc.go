package grpc

import (
	"context"

	"github.com/Mattiasg94/microservices-proto/golang/seatallocation"
	"github.com/Mattiasg94/microservices/seatallocation/internal/application/domain"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a Adapter) AllocateSeat(ctx context.Context, req *seatallocation.AllocateSeatRequest) (*emptypb.Empty, error) {
	err := a.api.AllocateSeat(ctx, domain.SeatAllocation{
		ID:         uuid.New().String(),
		FlightID:   req.FlightId,
		SeatNumber: req.SeatNumber,
	})

	return &emptypb.Empty{}, err
}

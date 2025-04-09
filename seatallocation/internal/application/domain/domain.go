package domain

type SeatAllocation struct {
	ID         string `json:"id"`
	FlightID   string `json:"flight_id"`
	SeatNumber string `json:"seat_number"`
}

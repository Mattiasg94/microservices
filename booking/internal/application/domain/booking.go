package domain

type Booking struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	FlightID   string `json:"flight_id"`
	SeatNumber string `json:"seat_number"`
}

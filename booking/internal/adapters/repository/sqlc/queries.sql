-- name: CreateBooking :one
INSERT INTO bookings (
  id,
  customer_id,
  flight_id,
  seat_number 
) VALUES (
  ?, ?, ?, ?
) RETURNING *;


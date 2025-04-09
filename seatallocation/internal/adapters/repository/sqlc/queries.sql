-- name: SaveSeatAllocation :exec
INSERT INTO seat_allocations (
  id,
  flight_id,
  seat_number
) VALUES (
  ?, ?, ?
);
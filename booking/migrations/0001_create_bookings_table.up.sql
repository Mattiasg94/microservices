CREATE TABLE bookings (
  id TEXT PRIMARY KEY,
  customer_id TEXT NOT NULL,
  flight_id TEXT NOT NULL,
  seat_number TEXT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
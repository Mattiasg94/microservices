package repository

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"

	sqlc "github.com/Mattiasg94/microservices/booking/internal/adapters/repository/sqlc"
	"github.com/Mattiasg94/microservices/booking/internal/application/domain"
	"github.com/google/uuid"
)

type Adapter struct {
	repo *sql.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	runMigrations(db, "../booking/migrations/0001_create_bookings_table.up.sql")

	return &Adapter{
		repo: db,
	}, nil
}

func (r *Adapter) Create(ctx context.Context, booking domain.Booking) (domain.Booking, error) {
	queries := sqlc.New(r.repo)

	params := sqlc.CreateBookingParams{
		ID:         uuid.New().String(),
		CustomerID: booking.CustomerID,
		FlightID:   booking.FlightID,
	}

	createdBooking, err := queries.CreateBooking(ctx, params)
	if err != nil {
		return domain.Booking{}, err
	}

	return domain.Booking{
		ID:         createdBooking.ID,
		CustomerID: createdBooking.CustomerID,
		FlightID:   createdBooking.FlightID,
	}, nil
}

// OBS: This is a temporary function to run migrations.
func runMigrations(db *sql.DB, migrationPath string) error {
	absPath, err := filepath.Abs(migrationPath)
	if err != nil {
		return err
	}

	sqlBytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))
	return err
}

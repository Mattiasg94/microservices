package repository

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log/slog"
	"path/filepath"

	"github.com/Mattiasg94/microservices/seatallocation/internal/adapters/repository/sqlc"
	"github.com/Mattiasg94/microservices/seatallocation/internal/application/domain"
	"github.com/google/uuid"
)

type Adapter struct {
	repo *sql.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		slog.Error("Failed to connect to SQLite database", "error", err)
	}
	runMigrations(db, "../seatallocation/migrations/0001_create_seat_allocation_table.up.sql")

	return &Adapter{
		repo: db,
	}, nil
}

func (r *Adapter) Save(ctx context.Context, seatAllocation domain.SeatAllocation) error {
	queries := sqlc.New(r.repo)

	params := sqlc.SaveSeatAllocationParams{
		ID:         uuid.New().String(),
		SeatNumber: seatAllocation.SeatNumber,
	}

	err := queries.SaveSeatAllocation(ctx, params)
	if err != nil {
		return err
	}

	return nil
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

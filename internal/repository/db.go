package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"task-management/internal/config"
)

func ConnectDB(cfg *config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}
	
	log.Println("Database connected successfully")

	return pool
}
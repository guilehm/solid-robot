package postgresStorage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func NewPostgresStorage() *Postgres {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/solid")

	// db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err.Error())
	}

	return &Postgres{
		DB: db,
	}
}

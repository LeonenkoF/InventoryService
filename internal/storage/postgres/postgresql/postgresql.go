package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn
}

func NewStorage() *Postgres {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:qwerty@localhost:5432/PC_inventory")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return &Postgres{conn: conn}
}

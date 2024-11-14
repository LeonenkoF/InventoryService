package tests

import (
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

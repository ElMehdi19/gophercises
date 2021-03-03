package db

import "database/sql"

// DB type
type DB struct {
	db *sql.DB
}

// Phone relation schema
type Phone struct {
	ID    int
	Value string
}

package db

import "database/sql"

type DB struct {
	*sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{db}
}

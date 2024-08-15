package handler

import "Auth-Server/internal/data/database"

type DbConfig struct {
	q *database.Queries
}

func NewDBConfig(q *database.Queries) *DbConfig {
	return &DbConfig{q: q}
}

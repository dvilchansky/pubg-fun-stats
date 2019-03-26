package repositories

import (
	"database/sql"
)

type MatchRepository interface {
}

type matchSQLRepository struct {
	DB *sql.DB
}

func NewMatchSQLRepository(DB *sql.DB) MatchRepository {
	return &matchSQLRepository{DB}
}

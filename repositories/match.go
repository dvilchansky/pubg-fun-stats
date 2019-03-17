package repositories

import (
	"database/sql"
	"fmt"
	"pubg-fun-stats/parser/models/match"
)

type MatchRepository interface {
	Fetch(lim int64) ([]*match.Match, error)
	FetchByIDs(matchIDs string) (map[string]*match.Match, error)
	Store(match *match.Match) error
}

type matchSQLRepository struct {
	Conn *sql.DB
}

func NewMatchSQLRepository(conn *sql.DB) MatchRepository {
	return &matchSQLRepository{conn}
}

func (msr *matchSQLRepository) Fetch(lim int64) ([]*match.Match, error) {
	query := "SELECT match_id, map_name, created_at, duration, game_mode FROM `match`"
	rows, err := msr.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*match.Match, 0)
	for rows.Next() {
		t := new(match.Match)
		err = rows.Scan(
			&t.ID,
			&t.MapName,
			&t.CreatedAt,
			&t.Duration,
			&t.GameMode,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (msr *matchSQLRepository) Store(m *match.Match) error {
	_, err := msr.Conn.Exec("INSERT IGNORE INTO `match` SET `id` = ?, `created_at` = ?, `duration` = ?, `game_mode` = ?, `map_name` = ?",
		m.ID, m.CreatedAt, m.Duration, m.GameMode, m.MapName)
	if err != nil {
		return err
	}
	return nil
}

func (msr *matchSQLRepository) FetchByIDs(matchIDs string) (map[string]*match.Match, error) {
	query := "SELECT * FROM `match` where id IN (%s)"
	rows, err := msr.Conn.Query(fmt.Sprintf(query, matchIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]*match.Match, 0)
	for rows.Next() {
		m := new(match.Match)
		err = rows.Scan(
			&m.ID,
			&m.CreatedAt,
			&m.Duration,
			&m.GameMode,
			&m.MapName,
		)
		if err != nil {
			return nil, err
		}
		result[m.ID] = m
	}
	return result, nil
}

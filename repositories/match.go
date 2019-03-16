package repositories

import (
	"database/sql"
	"fmt"
	"pubg-fun-stats/models"
)

type MatchRepository interface {
	Fetch(lim int64) ([]*models.Match, error)
	FetchByIDs(matchIDs string) (map[string]*models.Match, error)
	//GetByID(id string) (*models.Match, error)
	Store(match *models.Match) error
	//StoreMany(matches []*models.Match) (string, error)
}

type matchSQLRepository struct {
	Conn *sql.DB
}

func NewMatchSQLRepository(conn *sql.DB) MatchRepository {
	return &matchSQLRepository{conn}
}

func (msr *matchSQLRepository) Fetch(lim int64) ([]*models.Match, error) {
	query := "SELECT match_id, map_name, created_at, duration, game_mode FROM `match`"
	rows, err := msr.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Match, 0)
	for rows.Next() {
		t := new(models.Match)
		err = rows.Scan(
			&t.MatchID,
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

func (msr *matchSQLRepository) Store(match *models.Match) error {
	_, err := msr.Conn.Exec("INSERT IGNORE INTO `match` SET `match_id` = ?, `shard_id` = ?, `created_at` = ?, `duration` = ?, `game_mode` = ?, `map_name` = ?",
		match.MatchID, match.ShardID, match.CreatedAt, match.Duration, match.GameMode, match.MapName)
	if err != nil {
		panic(err.Error())
		return err
	}
	return nil
}

func (msr *matchSQLRepository) FetchByIDs(matchIDs string) (map[string]*models.Match, error) {
	query := "SELECT * FROM `match` where match_id IN (%s)"
	rows, err := msr.Conn.Query(fmt.Sprintf(query, matchIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]*models.Match, 0)
	for rows.Next() {
		m := new(models.Match)
		err = rows.Scan(
			&m.MatchID,
			&m.ShardID,
			&m.CreatedAt,
			&m.Duration,
			&m.GameMode,
			&m.MapName,
		)
		if err != nil {
			return nil, err
		}
		result[m.MatchID] = m
	}
	return result, nil
}

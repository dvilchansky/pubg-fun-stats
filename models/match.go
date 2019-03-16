package models

import "time"

type Match struct {
	MatchID   string    `json:"match_id"`
	ShardID   string    `json:"shard_id"`
	CreatedAt time.Time `json:"created_at"`
	Duration  int       `json:"duration"`
	GameMode  string    `json:"game_mode"`
	MapName   string    `json:"map_name"`
}

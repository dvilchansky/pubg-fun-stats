package models

import (
	"fmt"
	"github.com/dvilchansky/gopubg"
)

func InsertMatch(a *gopubg.API, matchID string) {
	match, err := a.RequestMatch("steam", matchID)
	if err != nil {
		panic(err.Error())
	}
	if match.MapName == "Range_Main" {
		return
	}
	insert, err := db.Connection.Query(
		"INSERT IGNORE INTO matches "+
			"SET match_id = ?,shard_id = ?,created_at = ?,duration = ?,game_mode = ?, map_name = ?",
		match.ID, match.ShardID, match.CreatedAt, match.Duration, match.GameMode, match.MapName)
	for _, roster := range match.Rosters {
		insert, err := db.Connection.Query(
			"INSERT IGNORE INTO rosters SET match_id = ?, roster_id = ?, shard_id = ?, rank = ?, team_id = ?, won = ?",
			match.ID, roster.ID, roster.ShardID, roster.Stats.Rank, roster.Stats.TeamID, roster.Won)
		if err != nil {
			panic(err.Error())
		}
		insert.Close()
		for _, p := range roster.Participants {
			insert, err := db.Connection.Query(
				"INSERT IGNORE INTO participants SET match_id = ?, roster_id = ?, participant_id = ?, assists = ?,"+
					"boosts = ?, damage_dealt = ?, death_type = ?, headshot_kills = ?, heals = ?, kill_place = ?, "+
					"kill_streaks = ?, kills = ?, longest_kill = ?, name = ?, player_id = ?, revives = ?,"+
					"ride_distance = ?,road_kills = ?, team_kills = ?, time_survived = ?, vehicle_destroys = ?, walk_distance = ?, win_place = ?, swim_distance = ?",
				match.ID, roster.ID, p.ID, p.Stats.Assists, p.Stats.Boosts, p.Stats.DamageDealt, p.Stats.DeathType, p.Stats.HeadshotKills,
				p.Stats.Heals, p.Stats.KillPlace, p.Stats.KillStreaks, p.Stats.Kills, p.Stats.LongestKill, p.Stats.Name, p.Stats.PlayerID,
				p.Stats.Revives, p.Stats.RideDistance, p.Stats.RoadKills, p.Stats.TeamKills, p.Stats.TimeSurvived, p.Stats.VehicleDestroys,
				p.Stats.WalkDistance, p.Stats.WinPlace, p.Stats.SwimDistance)
			if err != nil {
				panic(err.Error())
			}
			insert.Close()
		}
	}

	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

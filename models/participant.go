package models

type Participant struct {
	MatchID       string  `json:"match_id"`
	Name          string  `json:"name"`
	PlayerID      string  `json:"player_id"`
	ParticipantID string  `json:"participant_id"`
	RosterID      string  `json:"roster_id"`
	DamageDealt   float64 `json:"damage_dealt"`
	DeathType     string  `json:"death_type"`
	HeadshotKills int     `json:"headshot_kills"`
	Kills         int     `json:"kills"`
	LongestKill   float64 `json:"longest_kill"`
	TimeSurvived  float64 `json:"time_survived"`
	WinPlace      int     `json:"win_place"`
}

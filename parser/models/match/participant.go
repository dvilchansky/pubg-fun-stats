package match

// Participant represents a player for a match
type Participant struct {
	ID    string `jsonapi:"primary,participant"`
	Stats struct {
		Name         string  `jsonapi:"attr,name"`
		DamageDealt  float64 `jsonapi:"attr,damageDealt"`
		WinPlace     int     `jsonapi:"attr,winPlace"`
		Kills        int     `jsonapi:"attr,kills"`
		TimeSurvived float64 `jsonapi:"attr,timeSurvived"`
		//Actor   string `jsonapi:"attr,actor"`
		//ShardID string `jsonapi:"attr,shardId"`
		//Assists         int     `jsonapi:"attr,assists"`
		//Boosts          int     `jsonapi:"attr,boosts"`
		//DeathType       string  `jsonapi:"attr,deathType"`
		//HeadshotKills   int     `jsonapi:"attr,headshotKills"`
		//Heals           int     `jsonapi:"attr,heals"`
		//KillPlace       int     `jsonapi:"attr,killPlace"`
		//KillStreaks     int     `jsonapi:"attr,killStreaks"`
		//LongestKill     int     `jsonapi:"attr,longestKill"`
		//PlayerID        string  `jsonapi:"attr,playerId"`
		//Revives         int     `jsonapi:"attr,revives"`
		//RideDistance    float64 `jsonapi:"attr,rideDistance"`
		//RoadKills       int     `jsonapi:"attr,roadKills"`
		//TeamKills       int     `jsonapi:"attr,teamKills"`
		//VehicleDestroys int     `jsonapi:"attr,vehicleDestroys"`
		//WalkDistance    float64 `jsonapi:"attr,walkDistance"`
		//SwimDistance    float64 `jsonapi:"attr,swimDistance"`
	} `jsonapi:"attr,stats"`
}

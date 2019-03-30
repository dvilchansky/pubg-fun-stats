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
	} `jsonapi:"attr,stats"`
}

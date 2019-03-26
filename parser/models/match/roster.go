package match

// Roster represents a team of participants
type Roster struct {
	ID      string `jsonapi:"primary,roster"`
	ShardID string `jsonapi:"attr,shardId"`
	Stats   struct {
		Rank   int `jsonapi:"attr,rank"`
		TeamID int `jsonapi:"attr,teamId"`
	} `jsonapi:"attr,stats"`
	Won          string         `jsonapi:"attr,won"`
	Participants []*Participant `jsonapi:"relation,participants"`
}

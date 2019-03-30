package match

import (
	"github.com/google/jsonapi"
	"io"
	"time"
)

// Match structure represent data related to a PUBG match
type Match struct {
	ID           string    `jsonapi:"primary,match"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration     int       `jsonapi:"attr,duration"`
	GameMode     string    `jsonapi:"attr,gameMode"`
	MapName      string    `jsonapi:"attr,mapName"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	ShardID      string    `jsonapi:"attr,shardId"`
	TitleID      string    `jsonapi:"attr,titleId"`
	Rosters      []*Roster `jsonapi:"relation,rosters"`
	Assets       []*Assets `jsonapi:"relation,assets"`
}

type Assets struct {
	ID          string    `jsonapi:"primary,asset"`
	CreatedAt   time.Time `jsonapi:"attr,createdAt,iso8601"`
	URL         string    `jsonapi:"attr,URL"`
	Name        string    `jsonapi:"attr,name"`
	Description string    `jsonapi:"attr,description"`
}

// ParseMatch parses a json response containing matches information
func ParseMatch(in io.Reader) (*Match, error) {
	match := new(Match)
	err := jsonapi.UnmarshalPayload(in, match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

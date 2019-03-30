package player

import (
	"errors"
	"github.com/google/jsonapi"
	"io"
	"reflect"
	"time"
)

// Player structure represents a player entry
type Player struct {
	ID           string    `jsonapi:"primary,player"`
	Name         string    `jsonapi:"attr,name"`
	ShardID      string    `jsonapi:"attr,shardId"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	UpdatedAt    time.Time `jsonapi:"attr,updatedAt,iso8601"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	TitleID      string    `jsonapi:"attr,titleId"`
	Matches      []*Match  `jsonapi:"relation,matches"`
}

// Match structure represent data related to a PUBG match
type Match struct {
	ID string `jsonapi:"primary,match"`
}

// ParsePlayers parses a json response containing players information
func ParsePlayers(in io.Reader) ([]*Player, error) {
	result, err := jsonapi.UnmarshalManyPayload(in, reflect.TypeOf(new(Player)))
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(result))
	for idx, elt := range result {
		player, ok := elt.(*Player)
		if !ok {
			return nil, errors.New("failed to convert players")
		}
		players[idx] = player
	}

	return players, nil
}

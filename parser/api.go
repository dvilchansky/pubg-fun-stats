package gopubg

import (
	"fmt"
	"net/url"
	"pubg-fun-stats/parser/models/match"
	"pubg-fun-stats/parser/models/player"
	"pubg-fun-stats/parser/models/telemetry"
)

//API struct for holding the key
type API struct {
	Key string
}

//NewAPI creating a new API from a key
func NewAPI(key string) *API {
	return &API{
		Key: key,
	}
}

//RequestStatus A function that prints out the current status of the API Key
func (a *API) RequestStatus() error {
	endpointURL := "https://api.pubg.com/status"

	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return err
	}

	fmt.Printf("data:\n%s\n", buffer)

	return nil
}

//RequestPlayerByName A function that takes a player name, and returns either that players data, or an error
func (a *API) RequestPlayerByName(playerName string) (*player.Player, error) {
	parameters := url.Values{
		"filter[playerNames]": {playerName},
	}

	endpointURL := fmt.Sprintf("https://api.pubg.com/shards/steam/players?%s", parameters.Encode())

	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return nil, err
	}
	players, err := player.ParsePlayers(buffer)
	if err != nil {
		return nil, err
	}
	return players[0], nil
}

//RequestMatch given a shard and a match_id string will print either match info, or a error
func (a *API) RequestMatch(matchID string) (*match.Match, error) {
	endpointURL := fmt.Sprintf("https://api.pubg.com/shards/steam/matches/%s", matchID)
	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return nil, err
	}
	return match.ParseMatch(buffer)
}

//RequestMatch given a shard and a match_id string will print either match info, or a error
func (a *API) RequestTelemetry(endpointURL string) (*telemetry.Telemetry, error) {
	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return nil, err
	}
	return telemetry.ParseTelemetry(buffer)
}

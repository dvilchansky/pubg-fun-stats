package main

import (
	"github.com/dvilchansky/gopubg"
	"pubg-fun-stats/db"
	"pubg-fun-stats/models/dbmatches"
	"pubg-fun-stats/settings"
)

func main() {
	// Initialise
	db.GetConnection()
	api := gopubg.NewAPI(settings.API_KEY)
	// Request player by name
	players, err := api.RequestSinglePlayerByName("steam-eu", `Xrage12`)
	if err != nil {
		panic(err.Error())
	}
	for _, p := range players {
		for _, m := range p.Matches {
			dbmatches.InsertMatch(api, m.ID)
		}
	}
}

package main

import (
	"github.com/dvilchansky/gopubg"
	"pubg-fun-stats/settings"
)

func main() {
	api := gopubg.NewAPI(settings.Key)
	api.RequestSinglePlayerByName(settings.Region, settings.Username)
}

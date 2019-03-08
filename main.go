package main

import (
	"github.com/dvilchansky/gopubg"
	"pubg-fun-stats/db"
	"pubg-fun-stats/models/dbplayers"
	"pubg-fun-stats/settings"
)

func main() {
	db.GetConnection()
	api := gopubg.NewAPI(settings.API_KEY)

	players, err := api.RequestSinglePlayerByName("steam-eu", `n0ic3`)
	if err != nil {
		panic(err.Error())
	}

	err = dbplayers.Insert(players)
	if err != nil {
		panic(err.Error())
	}

	//player, err := dbplayers.SelectOne(3)
	//if err != nil {
	//	panic(err.Error())
	//}

	//pp.Println(player)
}

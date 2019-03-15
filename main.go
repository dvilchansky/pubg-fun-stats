package main

import (
	"github.com/dvilchansky/gopubg"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"pubg-fun-stats/controllers"
	"pubg-fun-stats/funstats"
	"pubg-fun-stats/settings"
	"runtime"
)

var (
	ConcurencyLevel = runtime.NumCPU() * 22
	API             = gopubg.NewAPI(settings.API_KEY)
)

func main() {
	r := iris.Default()
	r.Post("/players/:name", PlayersHandler)
	r.StaticWeb("/", "./public/dist")
	//ws := websocket.New(websocket.Config{})
	todosRouter := r.Party("/players/{name}", PlayersHandler)
	todosRouter.Any("/iris-ws.js", websocket.ClientHandler())
	r.Run(iris.Addr(":8080"))
}

func GetMatches(api *gopubg.API, playerName string) []*funstats.SQLMatch {
	players, err := api.RequestSinglePlayerByName("steam", playerName)
	if err != nil {
		panic(err.Error())
	}
	var response []*funstats.SQLMatch
	for _, p := range players {
		response = controllers.ProcessMatches(ConcurencyLevel, api, p.Matches)
	}
	return response
}

func PlayersHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"data": GetMatches(API, ctx.Params().Get("name")),
	})
}

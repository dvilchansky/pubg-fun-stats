package main

import (
	"database/sql"
	"github.com/dvilchansky/gopubg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"pubg-fun-stats/controllers"
	"pubg-fun-stats/repositories"
	"pubg-fun-stats/services"
	"pubg-fun-stats/settings"
)

var (
	API    = gopubg.NewAPI(settings.API_KEY)
	DBConn *sql.DB
)

func match(app *mvc.Application) {
	repo := repositories.NewMatchSQLRepository(DBConn)
	matchService := services.NewMatchService(repo, API)
	app.Register(matchService)
	app.Handle(new(controllers.MatchController))
}

//func GetMatches(api *gopubg.API, playerName string) []*funstats.SQLMatch {
//	player, err := api.RequestPlayerByName(playerName)
//	if err != nil {
//		panic(err.Error())
//	}
//	return controllers.ProcessMatches(ConcurencyLevel, api, player.Matches)
//}

func main() {
	app := iris.Default()
	var err error
	DBConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/pubg_fun_stats?parseTime=true")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = DBConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer DBConn.Close()
	mvc.Configure(app.Party("/api/players/{name}"), match)
	app.StaticWeb("/", "./public/dist")
	app.Run(iris.Addr(":8080"))
}

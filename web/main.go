package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"pubg-fun-stats/parser"
	"pubg-fun-stats/repositories"
	"pubg-fun-stats/settings"
	"pubg-fun-stats/web/controllers"
	"pubg-fun-stats/web/services"
)

var (
	API    = gopubg.NewAPI(settings.API_KEY)
	DBConn *sql.DB
)

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

// Match handler
func match(app *mvc.Application) {
	repo := repositories.NewMatchSQLRepository(DBConn)
	matchService := services.NewMatchService(repo, API)
	app.Register(matchService)
	app.Handle(new(controllers.MatchController))
}

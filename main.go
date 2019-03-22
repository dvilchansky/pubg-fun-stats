package main

import (
	"database/sql"
	"fmt"
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
	API = gopubg.NewAPI(settings.API_KEY)
	DB  *sql.DB
)

func main() {
	app := iris.Default()
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		settings.DB_USER, settings.DB_PASSWORD, settings.DB_HOST, settings.DB_PORT, settings.DB_NAME))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer DB.Close()
	mvc.Configure(app.Party("/api/players/{name}"), match)
	mvc.Configure(app.Party("/api/telemetry/"), telemetry)
	app.StaticWeb("/", "./web/public/dist")
	app.Run(iris.Addr(":8080"))
}

// Match handler
func match(app *mvc.Application) {
	matchService := services.NewMatchService(repositories.NewMatchSQLRepository(DB), API)
	app.Register(matchService)
	app.Handle(new(controllers.MatchController))
}

// Match handler
func telemetry(app *mvc.Application) {
	telemetryService := services.NewTelemetryService(repositories.NewTelemetrySQLRepository(DB), API)
	app.Register(telemetryService)
	app.Handle(new(controllers.TelemetryController))
}

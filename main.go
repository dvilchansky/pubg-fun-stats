package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"os"
	"pubg-fun-stats/parser"
	"pubg-fun-stats/repositories"
	"pubg-fun-stats/web/controllers"
	"pubg-fun-stats/web/services"
)

func init() {
	API = gopubg.NewAPI(os.Getenv(`PUBG_API_KEY`))
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv(`DB_USER`),
		os.Getenv(`DB_PASS`),
		os.Getenv(`DB_HOST`),
		os.Getenv(`DB_PORT`),
		os.Getenv(`DB_NAME`)))
	if err != nil {
		panic(err.Error())
	}
}

var (
	API *gopubg.API
	DB  *sql.DB
)

func main() {
	app := iris.Default()
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

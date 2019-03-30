package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"pubg-fun-stats/parser"
	"pubg-fun-stats/web/controllers"
	"pubg-fun-stats/web/services"
	"time"
)

var (
	API          *gopubg.API
	CacheHandler = cache.Handler(30 * time.Minute)
)

func main() {
	API = gopubg.NewAPI(os.Getenv(`PUBG_API_KEY`))
	app := iris.Default()
	mvc.Configure(app.Party(`/api/players/{name}`), match)
	mvc.Configure(app.Party(`/api/telemetry/`), telemetry)
	app.StaticWeb("/", `./web/public/dist`)
	err := app.Run(iris.Addr(`:` + os.Getenv(`PORT`)))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

// Match handler
func match(app *mvc.Application) {
	app.Router.Use(CacheHandler)
	matchService := services.NewMatchService(API)
	app.Register(matchService)
	app.Handle(new(controllers.MatchController))
}

// Match handler
func telemetry(app *mvc.Application) {
	app.Router.Use(CacheHandler)
	telemetryService := services.NewTelemetryService(API)
	app.Register(telemetryService)
	app.Handle(new(controllers.TelemetryController))
}

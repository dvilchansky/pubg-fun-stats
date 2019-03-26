package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/mvc"
	"github.com/spf13/viper"
	"pubg-fun-stats/parser"
	"pubg-fun-stats/repositories"
	"pubg-fun-stats/web/controllers"
	"pubg-fun-stats/web/services"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	API = gopubg.NewAPI(viper.GetString(`pubg-api.key`))
	//DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	//	viper.GetString(`database.user`),
	//	viper.GetString(`database.pass`),
	//	viper.GetString(`database.host`),
	//	viper.GetString(`database.port`),
	//	viper.GetString(`database.name`)))
}

var (
	API          *gopubg.API
	DB           *sql.DB
	CacheHandler = cache.Handler(30 * time.Minute)
)

func main() {
	app := iris.Default()
	//defer DB.Close()
	mvc.Configure(app.Party("/api/players/{name}"), match)
	mvc.Configure(app.Party("/api/telemetry/"), telemetry)
	app.StaticWeb("/", "./web/public/dist")
	app.Run(iris.Addr(":8080"))
}

// Match handler
func match(app *mvc.Application) {
	app.Router.Use(CacheHandler)
	matchService := services.NewMatchService(repositories.NewMatchSQLRepository(DB), API)
	app.Register(matchService)
	app.Handle(new(controllers.MatchController))
}

// Match handler
func telemetry(app *mvc.Application) {
	app.Router.Use(CacheHandler)
	telemetryService := services.NewTelemetryService(repositories.NewTelemetrySQLRepository(DB), API)
	app.Register(telemetryService)
	app.Handle(new(controllers.TelemetryController))
}

package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"pubg-fun-stats/app/settings"
)

var Connection *sql.DB

func GetConnection() (*sql.DB, error) {
	var err error
	Connection, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		settings.DB_USER, settings.DB_PASSWORD, settings.DB_HOST, settings.DB_PORT, settings.DB_NAME))
	if err != nil {
		panic(err.Error())
	}
	return Connection, err
}

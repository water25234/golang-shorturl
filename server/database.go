package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/water25234/golang-shorturl/config"
)

var DB *sql.DB
var err error

func InitDataBase() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetAppConfig().DBHost,
		config.GetAppConfig().DBPort,
		config.GetAppConfig().DBUsername,
		config.GetAppConfig().DBPassword,
		config.GetAppConfig().DBDatabase,
	)

	DB, err = sql.Open(config.GetAppConfig().DBConnection, psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

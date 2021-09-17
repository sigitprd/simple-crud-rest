package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sigitprd/simple-crud-rest/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
}

func CreateConnection() *sql.DB {
	return db
}

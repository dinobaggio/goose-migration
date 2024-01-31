package config

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

type DBConfig struct {
	connection string
	hostname   string
	port       string
	user       string
	password   string
	database   string
}

func init() {
	var dbConfig DBConfig

	dbConfig.hostname = "127.0.0.1"
	dbConfig.port = "33007"
	dbConfig.user = "root"
	dbConfig.password = "password"
	dbConfig.database = "goose-migration"
	dbConfig.connection = "mysql"

	consStr := dbConfig.user + ":" + dbConfig.password + "@tcp(" + dbConfig.hostname + ":" + dbConfig.port + ")/" + dbConfig.database + "?parseTime=true"
	dsn := flag.String("dsn", consStr, "connection data source name")
	flag.Parse()

	if len(*dsn) == 0 {
		log.Fatal("missing dsn flag")
	}
	var err error
	pool, err = sql.Open(dbConfig.connection, *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
}

func SQLDBConn() *sql.DB {
	return pool
}

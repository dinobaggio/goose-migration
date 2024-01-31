package main

import (
	"goose-migration/config"

	"github.com/pressly/goose"
)

func main() {
	sqlDB := config.SQLDBConn()

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	err := goose.Up(sqlDB, "migrations")
	// err := goose.Down(sqlDB, "migrations")
	if err != nil {
		panic(err)
	}
}

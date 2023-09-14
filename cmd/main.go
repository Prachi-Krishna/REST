package main

import (
	"rest/config"
	"rest/db"

	"rest/pkg/api/seed"
	"rest/pkg/server"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		seed.Migrate(dbHandler)
	}

	server.Setup(dbHandler)
}

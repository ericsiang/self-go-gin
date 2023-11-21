package main

import (
	"api/database/migrate"
	"api/initialize"
	"api/router"
	"strconv"
)

func main() {
	initialize.InitConfig()
	migrate.Migrate() // migrate database

	// Set Router
	r := router.Router()

	// Listen and Server
	serverPort := ":" + strconv.Itoa(initialize.ServerEnv.GetServerPort())
	r.Run(serverPort)
}

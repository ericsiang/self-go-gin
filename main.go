package main

import (
	"api/database/migrate"
	"api/initialize"
	"api/model"

	// "go.uber.org/zap"
	// "strconv"
	// "api/router"
)

func main() {
	initialize.InitConfig()
	migrate.Migrate()
	use := model.Users{}
	use.GetUsersById(1)
	// zap.S().Info("user : ", user)
	// zap.S().Error("err : ", err)
	// r := router.Router()

	// // Listen and Server
	// serverPort := ":" + strconv.Itoa(initialize.ServerEnv.GetServerPort())

	// r.Run(serverPort)
}

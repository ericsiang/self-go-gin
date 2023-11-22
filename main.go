package main

import (
	"api/database/migrate"
	"api/initialize"
	"api/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var quit chan os.Signal

func main() {
	initialize.InitSetting()
	migrate.Migrate() // migrate database
	quit = make(chan os.Signal, 1)
	// Set Router
	router := router.Router(quit)
	// Listen and Server
	// serverPort := ":" + strconv.Itoa(initialize.GetServerEnv().GetServerPort())
	// r.Run(serverPort)

	//優雅的關閉服務(服務端關機命令發出後不是立即關機)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(initialize.GetServerEnv().GetServerPort()),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Error("Server listen : %s\n", err)
		}else{
			zap.S().Info("Server listening ......")
		}
	}()

	
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Preparing Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Error("Shutdown Error :", err)
	}else{
		zap.S().Info("Server exiting")
	}
	
}

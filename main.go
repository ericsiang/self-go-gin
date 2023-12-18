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
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var wg sync.WaitGroup

func main() {
	initialize.InitSetting()
	migrate.Migrate() // migrate database
	httpServerRun()

	//測試 log 切割
	// for i := 0; i < 2000; i++ {
	// 	wg.Add(2)
	// 	go simpleHttpGet("www.baidu.com")
	// 	go simpleHttpGet("https://www.baidu.com")
	// }
	// wg.Wait()
}

func httpServerRun() {
	quit := make(chan os.Signal, 1)
	// Set Router
	router := router.Router(quit)
	// Listen and Server
	// serverPort := ":" + strconv.Itoa(initialize.GetServerEnv().GetServerPort())
	// r.Run(serverPort)

	//優雅的關閉服務(服務端關機命令發出後不會立即關機)
	//建立一個http.Server
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(initialize.GetServerEnv().GetServerPort()),
		Handler: router,
	}

	go func() {
		//啟動 http.Server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Error("Server listen : %s\n", err)
		}
	}()

	/*
		監聽等待 SIGINT 或 SIGTERM 信号
		SIGINT -> 由使用者在終端中按下 Ctrl+C 產生，用於請求進程中斷
		SIGTERM -> 系統預設的終止信號，當你使用 kill 命令（不帶任何信號選項）
	*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Preparing Shutdown Server ...")
	//建立超時上下文，Shutdown可以讓未處理的連線在這個時間內關閉
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//停止HTTP服务器
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Error("Shutdown Error :", err)
	}

	zap.S().Info("Server exiting")
}


package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"self_go_gin/gin_application/router"
	validlang "self_go_gin/gin_application/validate_lang"

	"self_go_gin/infra/env"
	"self_go_gin/infra/orm/gorm_mysql"
	"self_go_gin/util/jwt_secret"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	serverEnv = env.NewServerConfig()
)

// @title  Self go gin Swagger API
// @version 1.0
// @description swagger first example
// @host localhost:5000
// @accept 		json
// @schemes 	http https
// @securityDefinitions.apikey	JwtTokenAuth
// @in			header
// @name   		Authorization
// @description Use Bearer JWT Token
func main() {
	initSetting()
	httpServerRun()
}

func httpServerRun() {
	quit := make(chan os.Signal, 1)
	// Set Router
	router := router.Router(quit)
	// Listen and Server
	// serverPort := ":" + strconv.Itoa(initialize.GetServerEnv().GetServerPort())

	//優雅的關閉服務(服務端關機命令發出後不會立即關機)
	//建立一個http.Server
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(GetServerEnv().Port),
		Handler: router,
	}

	go func() {
		//啟動 http.Server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintln(os.Stderr, "Server listen : ", err)
			return
		}
	}()

	/*
		監聽等待 SIGINT 或 SIGTERM 信号
		SIGINT -> 由使用者在終端中按下 Ctrl+C 產生，用於請求進程中斷
		SIGTERM -> 系統預設的終止信號，當你使用 kill 命令（不帶任何信號選項）
	*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Preparing Shutdown Server ...")
	//建立超時上下文，Shutdown可以讓未處理的連線在這個時間內關閉
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//停止HTTP服务器
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Shutdown Error :", err)
		return
	}

	fmt.Println("Server exiting")
}

func initSetting() {
	env.InitEnv("../../conf/", serverEnv, initSetting)
	fmt.Printf("配置信息 : %+v\n", serverEnv)
	gin.SetMode(serverEnv.AppMode)
	gorm_mysql.InitMysql(GetServerEnv)
	// redis.InitRedis(GetServerEnv)
	jwt_secret.SetJwtSecret(GetServerEnv().JwtSecret)
	// vaildate 中文化
	if err := validlang.InitValidateLang("zh"); err != nil {
		fmt.Fprintln(os.Stderr, "init trans failed, err:", err)
		panic(err)
	}
}

// GetServerEnv 獲取服務配置
func GetServerEnv() *env.ServerConfig {
	return serverEnv
}

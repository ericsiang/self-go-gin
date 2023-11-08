package main

import (
	"api/initialize"
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitConfig()
	r := gin.New()
	logger := initialize.GetLogger()
	/* Add a ginzap middleware, which:
	 * - Logs all requests, like a combined access and error log.
	 * - Logs to stdout.
	 * - RFC3339 with UTC time format.
	 */
	r.Use(ginzap.Ginzap(logger, "", true))

	/* Logs all panic to error log
	 *  - stack means whether output the stack info.
	 */
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	//記錄 info 級別日志
	// logger.Info("info 級別日志")

	// //記錄 error 級別日志
	// logger.Error("error 級別日志")

	// logger.Fatal("fatal 級別日志")
	// logger.Warn("warn 級別日志")

	// Listen and Server
	r.Run(":8888")
}

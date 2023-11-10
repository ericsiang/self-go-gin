package router

import (
	"api/initialize"
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	switch initialize.ServerEnv.GetServerAppMode() {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
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

	router := gin.Default()
	v1 := router.Group("v1")

	// Example ping request.
	v1.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	v1.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	// logger.Info("info 級別日志")
	// logger.Error("error 級別日志")
	// logger.Fatal("fatal 級別日志")
	// logger.Warn("warn 級別日志"
	
	return router
}

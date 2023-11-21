package router

import (
	v1 "api/api/v1"
	"api/initialize"
	"api/middleware"
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
	logger := initialize.Logger
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

	//==============================   no auth group   =================================

	apiV1Group := router.Group("/api/v1")
	apiV1UsersGroup := apiV1Group.Group("/users")
	apiV1Group.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	CreateUser(apiV1UsersGroup)
	Login(apiV1UsersGroup)

	//==============================   auth group   =================================
	apiV1AuthGroup := router.Group("/api/v1/auth")
	apiV1AuthUsersGroup := router.Group("/api/v1/auth")
	apiV1AuthGroup.Use(middleware.JwtAuthMiddleware())
	{
		// Users
		Users(apiV1AuthUsersGroup)
	}

	// Example when panic happen.
	// apiV1Group.GET("/panic", func(c *gin.Context) {
	// 	panic("An unexpected error happen!")
	// })

	// logger.Info("info 級別日志")
	// logger.Error("error 級別日志")
	// logger.Fatal("fatal 級別日志")
	// logger.Warn("warn 級別日志"

	return router
}

// =================================   no auth group   =====================================
func CreateUser(router *gin.RouterGroup){
	router.POST("/", v1.CreateUser)
}	
func Login(router *gin.RouterGroup) {
	router.POST("/login", v1.UserLogin)
}

//=================================   auth group   =====================================
func Users(router *gin.RouterGroup) {
	router.GET("/users", v1.GetUsers)
	router.GET("/users/:id", v1.GetUsersById)
}
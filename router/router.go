package router

import (
	v1 "api/api/v1"
	"api/initialize"
	"api/middleware"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Router(quit chan os.Signal) *gin.Engine {
	switch initialize.GetServerEnv().GetServerAppMode() {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	logger := initialize.GetZapLogger()
	/* Add a ginzap middleware, which:
	 * - Logs all requests, like a combined access and error log.
	 * - Logs to stdout.
	 * - RFC3339 with UTC time format.
	 */
	router.Use(ginzap.Ginzap(logger, "", true))

	/* Logs all panic to error log
	 *  - stack means whether output the stack info.
	 */
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(cors.Default()) //跨域請求的中間件

	//==============================   no auth group   =================================

	apiV1Group := router.Group("/api/v1")
	apiV1UsersGroup := apiV1Group.Group("/users")
	apiV1Group.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	CreateUser(apiV1UsersGroup)
	Login(apiV1UsersGroup)

	//==============================   auth group   =================================
	apiV1AuthGroup := apiV1Group.Group("/auth")
	apiV1AuthUsersGroup := apiV1AuthGroup.Group("/users")
	apiV1AuthUsersGroup.Use(middleware.JwtAuthMiddleware())
	{
		// Users
		Users(apiV1AuthUsersGroup)
	}
	apiV1AuthAdminsGroup := apiV1AuthGroup.Group("/admins")
	Shutdown(apiV1AuthAdminsGroup,quit)

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
func CreateUser(router *gin.RouterGroup) {
	router.POST("/", v1.CreateUser)
}
func Login(router *gin.RouterGroup) {
	router.POST("/login", v1.UserLogin)
}

// =================================   auth group   =====================================
func Users(router *gin.RouterGroup) {
	router.GET("/", v1.GetUsers)
	router.GET("/:filterUsersId", v1.GetUsersById)
}


func Shutdown(router *gin.RouterGroup ,quit chan os.Signal) {
	router.GET("/shutdown", func(c *gin.Context) {
		close(quit)
		c.String(200, "shutdown")
	})
}
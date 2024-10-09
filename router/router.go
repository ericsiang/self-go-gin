package router

import (
	v1 "api/api/v1"
	"api/initialize"
	"api/middleware"
	"fmt"
	"os"

	// "strconv"
	"syscall"
	"time"

	_ "api/docs"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func setDefaultMiddlewares(router *gin.Engine) {
	zapLogger := initialize.GetZapLogger()
	/* Add a ginzap middleware, which:
	 * - Logs all requests, like a combined access and error log.
	 * - Logs to stdout.
	 * - RFC3339 with UTC time format.
	 */
	router.Use(ginzap.Ginzap(zapLogger, "", true))

	/* Logs all panic to error log
	 *  - stack means whether output the stack info.
	 */
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	})) //跨域請求的中間件
}

func Router(quit chan os.Signal) *gin.Engine {
	router := gin.New()
	setDefaultMiddlewares(router)
	registerSwagger(router)
	apiV1Group := router.Group("/api/v1")
	setNoAuthRoutes(apiV1Group)
	setAuthRoutes(apiV1Group, quit)
	return router
}

func registerSwagger(router *gin.Engine) {
	if gin.Mode() != gin.ReleaseMode {
		router.GET("/swagger-test/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setNoAuthRoutes(apiV1Group *gin.RouterGroup) {
	apiV1UsersGroup := apiV1Group.Group("/users")
	apiV1AdminsGroup := apiV1Group.Group("/admins")

	apiV1Group.Use(middleware.RateLimit("test-limit")).GET("/limit_ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})
	apiV1Group.Use(middleware.OpaMiddleware()).GET("/opa_ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	apiV1Group.GET("/logtest", func(c *gin.Context) {
		test := true
		if test {
			zap.L().Info("Logger  Success..",
				zap.String("GgGGGG", "200"))
		} else {
			zap.L().Error(
				"Logger  Error ..",
				zap.String("test", "just for test"))
		}
	})
	Login(apiV1UsersGroup, apiV1AdminsGroup)

}

func setAuthRoutes(apiV1Group *gin.RouterGroup, quit chan os.Signal) {
	apiV1AuthGroup := apiV1Group.Group("/auth")
	apiV1AuthGroup.Use(middleware.JwtAuthMiddleware())

	// Users
	apiV1AuthUsersGroup := apiV1AuthGroup.Group("/users")
	Users(apiV1AuthUsersGroup)

	// Admins
	apiV1AuthAdminsGroup := apiV1AuthGroup.Group("/admins")
	Admins(apiV1AuthAdminsGroup, quit)

}

// =================================   no auth group   =====================================

func Login(userRouter, adminRouter *gin.RouterGroup) {
	userRouter.POST("/login", v1.UserLogin)
	adminRouter.POST("/login", v1.AdminLogin)
}

// =================================   auth group   =====================================
func Users(router *gin.RouterGroup) {
	router.POST("", v1.CreateUser)
	router.GET("/:filterUsersId", v1.GetUsersById)
}

func Admins(router *gin.RouterGroup, quit chan os.Signal) {
	router.GET("/:filterAdminsId", v1.GetAdminsById)
	Shutdown(router, quit)
}

func Shutdown(router *gin.RouterGroup, quit chan os.Signal) {
	router.GET("/shutdown", func(c *gin.Context) {
		quit <- syscall.SIGINT
		c.String(200, "shutdown")
	})
}

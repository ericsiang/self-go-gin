package middleware

import (
	"self_go_gin/initialize"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

// 單一 server 使用
func RateLimit(redis_limit_key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := redis_rate.NewLimiter(initialize.GetRedisClient())
		// 限制每秒 5 個 request
		res, err := limiter.Allow(c, redis_limit_key, redis_rate.PerSecond(5))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}

		if res.Allowed == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": "Too many requests",
			})
			return
		}

		c.Next()
	}

}

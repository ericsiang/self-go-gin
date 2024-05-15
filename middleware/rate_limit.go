package middleware

import (
	"api/initialize"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RateLimit(redis_limit_key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := redis_rate.NewLimiter(initialize.GetRedisClient())
		res, err := limiter.Allow(c, redis_limit_key, redis_rate.PerMinute(10))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		}

		if res.Allowed == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": "Too many requests",
			})
		}

		c.Next()
	}

}

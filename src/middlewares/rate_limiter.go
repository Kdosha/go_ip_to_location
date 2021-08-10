package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"learning_go/configs"
	"log"
	"net/http"
	"time"
)

var LimiterTimeDuration = time.Duration(configs.LimiterTimeFrame) * time.Second
var cacheInstance = cache.New(LimiterTimeDuration, LimiterTimeDuration)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentCreditsUsed, _ := cacheInstance.Get(c.FullPath())
		if (currentCreditsUsed == nil) || (currentCreditsUsed == 0) {
			cacheInstance.SetDefault(c.FullPath(), 0)
			currentCreditsUsed = 0
		}

		if currentCreditsUsed.(int) > configs.LimiterMaxCredits-1 {
			c.JSON(http.StatusTooManyRequests, map[string]string{"error": "too many requests"})
			c.Abort()

		} else {
			c.Next()
			// after request
			err := cacheInstance.Increment(c.FullPath(), 1)
			if err != nil {
				log.Print("Cannot set data in cache")
			}
		}
	}
}

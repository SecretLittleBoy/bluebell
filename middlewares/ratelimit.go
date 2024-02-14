package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	rateLimitTokenBucket "github.com/juju/ratelimit" // 令牌桶
	rateLimitLeakyBucket "go.uber.org/ratelimit"     // 漏桶
)

// 基于令牌桶的限流中间件
// eg:RateLimiterTokenBucket(2*time.Second, 1)
func RateLimiterTokenBucket(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := rateLimitTokenBucket.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) != 1 {
			// 如果取不到令牌就返回响应
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}

// 基于漏桶的限流中间件
// eg:RateLimiterLeakyBucket(100)
func RateLimiterLeakyBucket(rate int) func(ctx *gin.Context) {
	// 生成一个限流器，
	rl := rateLimitLeakyBucket.New(rate)
	return func(c *gin.Context) {
		// 取水滴
		if time.Until(rl.Take()) > 0 {
			//time.Sleep(time.Until(rl.Take())) // 需要等这么长时间下一滴水才会滴下来
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

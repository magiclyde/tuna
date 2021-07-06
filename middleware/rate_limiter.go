/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午10:41
 * @note:
 * @refer: https://blog.jln.co/Rate-limit-with-Go-and-Gin/
 */

package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"sync"
	"time"
)

type RateKeyFunc func(ctx *gin.Context) (string, error)

type RateLimiter struct {
	fillInterval time.Duration
	capacity     int64
	rateKeygen   RateKeyFunc
	safeLimiters sync.Map
}

func (r *RateLimiter) get(ctx *gin.Context) (*ratelimit.Bucket, error) {
	key, err := r.rateKeygen(ctx)
	if err != nil {
		return nil, err
	}

	val, ok := r.safeLimiters.Load(key)
	if !ok {
		limiter := ratelimit.NewBucketWithQuantum(r.fillInterval, r.capacity, r.capacity)
		r.safeLimiters.Store(key, limiter)
		return limiter, nil
	}

	limiter, ok := val.(*ratelimit.Bucket)
	if !ok {
		return nil, errors.New("[middleware.ratelimit] fail to load limiter")
	}

	return limiter, nil
}

func (r *RateLimiter) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter, err := r.get(ctx)
		if err != nil || limiter.TakeAvailable(1) == 0 {
			if err == nil {
				err = errors.New("[middleware.ratelimit] too many requests")
			}
			ctx.AbortWithError(429, err)
			return
		} else {
			ctx.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
			ctx.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
			ctx.Next()
		}
	}
}

func NewRateLimiter(interval time.Duration, capacity int64, keyGen RateKeyFunc) *RateLimiter {
	return &RateLimiter{
		interval,
		capacity,
		keyGen,
		sync.Map{},
	}
}

// Allow only 60 requests per minute per ip
var DefaultIpLimiter = NewRateLimiter(time.Minute, 60, func(ctx *gin.Context) (string, error) {
	key := ctx.ClientIP()
	if key != "" {
		return key, nil
	}
	return "", errors.New("ip is missing")
})

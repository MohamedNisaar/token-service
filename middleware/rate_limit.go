package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
	"time"
)

var redisClient = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ip := strings.Split(r.RemoteAddr, ":")[0]
		key := "rate_limit" + ip

		count, _ := redisClient.Incr(ctx, key).Result()
		if count == 1 {
			redisClient.Expire(ctx, key, time.Minute)
		}

		if count > 10 {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
		}
		next.ServeHTTP(w, r)
	})
}

package http_server

import (
	"github.com/KartoonYoko/problem-details-storage/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func middlewareLogRequestDuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		logger.Log.Info(
			"request duration",
			zap.Int64("duration ms", duration.Milliseconds()),
		)
	})
}

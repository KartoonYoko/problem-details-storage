package http_server

import (
	"github.com/KartoonYoko/problem-details-storage/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

func middlewareLogRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		logger.Log.Info(
			"request info",
			zap.String("uri", r.RequestURI),
			zap.String("method", r.Method),
		)
	})
}

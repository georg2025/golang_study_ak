package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func main() {
	r := makeRouter()
	http.ListenAndServe("localhost:8080", r)
}

func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(LoggerMiddleware)
	r.Get("/1", firstAnswer)
	r.Get("/2", secondAnswer)
	r.Get("/3", thirdAnswer)
	r.NotFound(usualAnswer)
	return r
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, _ := zap.NewProduction()
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			duration := time.Since(start)
			logger.Info("handled request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", ww.Status()),
				zap.Duration("duration", duration),
				zap.String("remote_addr", r.RemoteAddr),
				zap.Int("response_size", ww.BytesWritten()),
			)
		}()
		next.ServeHTTP(ww, r)
	})
}

func firstAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func secondAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 2"))
}

func thirdAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 3"))
}

func usualAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

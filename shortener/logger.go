package shortener

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewStatusResponseWriter returns pointer to a new statusResponseWriter object
func NewStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader assigns status code and header to ResponseWriter of statusResponseWriter object
func (sw *statusResponseWriter) WriteHeader(statusCode int) {
	sw.statusCode = statusCode
	sw.ResponseWriter.WriteHeader(statusCode)
}

func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sw := NewStatusResponseWriter(w)
			defer func() {
				log.Printf(
					"[%s] [%d] (%v) %s %s",
					r.Method,
					sw.statusCode,
					time.Since(start),
					r.Host,
					r.URL.Path,
				)
			}()

			h.ServeHTTP(sw, r)
		})
	}
}

package logging

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

// RequestLogger returns a middleware that logs request start and request end
// with zerolog logger from the request context.
func RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()

			logRequestStart(r)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				elapsed := time.Since(t1)
				logRequestEnd(r, ww.Status(), ww.BytesWritten(), elapsed)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

func logRequestStart(r *http.Request) {
	hlog.FromRequest(r).Info().
		Msg("Request started")
}

func logRequestEnd(r *http.Request, status, bytes int, elapsed time.Duration) {
	elapsedMillis := float64(elapsed.Nanoseconds()) / 1000000

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	hlog.FromRequest(r).Info().
		Int("status_code", status).
		Str("http_method", r.Method).
		Str("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)).
		Str("remote_addr", r.RemoteAddr).
		Str("user_agent", r.UserAgent()).
		Int("bytes", bytes).
		Float64("elapsed_ms", elapsedMillis).
		Msg("Request completed")
}

package routes

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/mahendrafathan/go-cake/common/logger"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

func LoggingMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cid := xid.New().String()
			ctx := context.WithValue(r.Context(), logger.XCorrelationID, cid)

			startTime := time.Now()
			logRespWriter := NewLogResponseWriter(w)
			next.ServeHTTP(logRespWriter, r.WithContext(ctx))

			log.WithFields(log.Fields{
				logger.CORRELATION_ID: cid,
				"status":              logRespWriter.statusCode,
				"method":              r.Method,
				"path":                r.URL.EscapedPath(),
				"duration":            time.Since(startTime),
			}).Info("Incomming request")
		})
	}
}

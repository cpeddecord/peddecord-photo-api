package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func WithLogging(h http.Handler) http.Handler {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.WithFields(log.Fields{
			"url":     r.URL.Path,
			"method":  r.Method,
			"headers": r.Header,
		})

		start := time.Now()

		lrw := newLoggingResponseWriter(w)
		h.ServeHTTP(lrw, r)

		logger.WithFields(log.Fields{
			"status":   lrw.statusCode,
			"duration": (time.Since(start).Seconds() * 1000),
		}).Info("request end")
	})
}

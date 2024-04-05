package logger

import (
	"github.com/rs/zerolog/log"
)

// APILogger is a logger for API requests.
type APILogger struct {
}

// NewAPILogger creates a new APILogger.
func NewAPILogger() *APILogger {
	return &APILogger{}
}

// Info logs an info message.
func (a *APILogger) Info(method, route, msg string) *APILogger {
	log.
		Info().
		Str("method", method).
		Str("route", route).
		Msg(msg)

	return a
}

// Debug logs a debug message.
func (a *APILogger) Debug(method, route, msg string) *APILogger {
	log.
		Debug().
		Str("method", method).
		Str("route", route).
		Msg(msg)

	return a
}

// Err logs an error message.
func (a *APILogger) Err(method, route, msg string, statusCode int, err error) *APILogger {
	log.
		Err(err).
		Int("status", statusCode).
		Str("method", method).
		Str("route", route).
		Msgf(msg)

	return a
}

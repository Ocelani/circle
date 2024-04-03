package main

import "github.com/rs/zerolog/log"

// APILogger is a logger for API requests.
type APILogger struct {
	Method string
	Route  string
}

// NewAPILogger creates a new APILogger.
func NewAPILogger(method, route string) *APILogger {
	return &APILogger{
		Method: method,
		Route:  route,
	}
}

// Info logs an info message.
func (a *APILogger) Info(msg string) *APILogger {
	log.
		Info().
		Str("method", a.Method).
		Str("route", a.Route).
		Msg(msg)

	return a
}

// Debug logs a debug message.
func (a *APILogger) Debug(msg string) *APILogger {
	log.
		Debug().
		Str("method", a.Method).
		Str("route", a.Route).
		Msg(msg)

	return a
}

// Err logs an error message.
func (a *APILogger) Err(msg string, statusCode int, err error) *APILogger {
	log.
		Err(err).
		Int("status", statusCode).
		Str("method", a.Method).
		Str("route", a.Route).
		Msgf(msg)

	return a
}

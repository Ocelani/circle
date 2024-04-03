package main

import "github.com/rs/zerolog/log"

type APILogger struct {
	Method string
	Route  string
}

func NewAPILogger(method, route string) *APILogger {
	return &APILogger{
		Method: method,
		Route:  route,
	}
}

func (a *APILogger) Info(msg string) *APILogger {
	log.
		Info().
		Str("method", a.Method).
		Str("route", a.Route).
		Msg(msg)

	return a
}

func (a *APILogger) Debug(msg string) *APILogger {
	log.
		Debug().
		Str("method", a.Method).
		Str("route", a.Route).
		Msg(msg)

	return a
}

func (a *APILogger) Err(msg string, statusCode int, err error) *APILogger {
	log.
		Err(err).
		Int("status", statusCode).
		Str("method", a.Method).
		Str("route", a.Route).
		Msgf(msg)

	return a
}

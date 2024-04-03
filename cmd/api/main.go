package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// logLevel is the global log level for the application.
const logLevel = zerolog.DebugLevel

// init initializes the logger.
func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(logLevel)
	log.Debug().Msg("init")
}

// main is the entrypoint for the application.
func main() {
	log.Info().Msg("running...")

	server := http.NewServeMux()
	server.HandleFunc("/tb01", postTB01)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}

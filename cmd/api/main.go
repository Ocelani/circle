package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const logLevel = zerolog.DebugLevel

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(logLevel)
	log.Debug().Msg("init")
}

func main() {
	log.Info().Msg("running...")

	server := http.NewServeMux()
	server.HandleFunc("/tb01", postTB01)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}

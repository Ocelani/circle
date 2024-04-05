package main

import (
	"circle/cmd/api/internal"
	"circle/pkg/config"
	"circle/pkg/database"
	"circle/pkg/logger"
	"flag"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DefaultPort = "3000"             // DefaultPort is the default port for the server.
	LogLevel    = zerolog.DebugLevel // logLevel is the global log level for the application.
)

// init initializes the logger.
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	logger.Setup(LogLevel)
}

// main is the entrypoint for the application.
func main() {
	var sqlFile string
	flag.StringVar(&sqlFile, "sql", "", "sql file path to be executed")
	flag.Parse()

	log.Info().Msg("starting...")

	db, err := database.NewPostgreSQL(config.NewDatabase())
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}

	if sqlFile != "" {
		buf, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Panic().Err(err).Msg("failed to read SQL file")
		}
		go db.Exec(string(buf))
	}

	rep := internal.NewTB01Repository(db)
	svc := internal.NewTB01Service(rep)
	ctr := internal.NewTB01Controller(svc, rep, logger.NewAPILogger())
	app := http.NewServeMux()

	app.HandleFunc("/tb01", ctr.Post)

	port := GetPort()
	log.Info().Str("port", port).Msg("running...")

	if err := http.ListenAndServe(port, app); err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}

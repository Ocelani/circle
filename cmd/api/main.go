package main

import (
	"circle/cmd/api/internal"
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

// init initializes env vars and logger.
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	logger.Setup(LogLevel)
}

// main is the entrypoint for the application.
func main() {
	// parse flags
	var sqlFile string
	flag.StringVar(&sqlFile, "sql", "", "sql file path to be executed")
	flag.Parse()

	log.Info().Msg("api starting...")

	// open database connection
	db, err := database.NewPostgreSQL(NewDatabaseConfig())
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}

	// read and load sql file
	if sqlFile != "" {
		buf, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Panic().Err(err).Msg("failed to read SQL file")
		}
		go db.Exec(string(buf))
	}

	// create repository, service, controller and app
	rep := internal.NewTB01Repository(db)
	svc := internal.NewTB01Service(rep)
	ctr := internal.NewTB01Controller(svc, logger.NewAPILogger())
	app := http.NewServeMux()

	// routes
	app.HandleFunc("/tb01", ctr.Post)

	port := GetPort()
	log.Info().Str("port", port).Msg("running...")

	// start server
	if err := http.ListenAndServe(port, app); err != nil {
		log.Panic().Err(err).Msg("failed to start server")
	}
}

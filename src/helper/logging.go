package helper

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger *zerolog.Logger

func LoggingFile() {
	file, err := os.Create("task_logging.log")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create file")
		defer file.Close()
	}

	logger := zerolog.New(file).With().Timestamp().Logger()

	multi := zerolog.MultiLevelWriter(file, zerolog.ConsoleWriter{Out: os.Stdout})
	logger = logger.Output(multi)

	Logger = &logger

	// Example log statements
	// logger.Info().Str("module", "main").Msg("Starting the application")
	// logger.Warn().Str("module", "main").Msg("Something doesn't seem right")
	// logger.Error().Str("module", "main").Msg("An error occurred")

	// Wait for the logs to be flushed to the file
	time.Sleep(time.Second)

	// logger.Info().Str("module", "main").Msg("Exiting the application")
}

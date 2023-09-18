package main

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const path, name = "logs", "go_logs.log"

func SetGinLog() {
	gin.SetMode(gin.ReleaseMode)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logFile, err := os.Create(filepath.Join(path, name))

	if err != nil {
		log.Panic().Err(err)
	}

	gin.DefaultWriter = io.MultiWriter(logFile)
}

func ConfigLogger() {

	consoleOut := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RubyDate,
	}

	log.Logger = log.Output(consoleOut).With().Timestamp().Logger()
}

package main

import (
	"fmt"
	"os"

	"github.com/Freshify-HHS-ICT/backend/config"
	"github.com/MonkyMars/gecho"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(0)
	}
}

func run() error {
	// Load enviroment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: could not load .env file")
	}

	cfg := config.GetConfig()

	logLevel := gecho.ParseLogLevel(cfg.Server.LogLevel)
	logger := gecho.NewLogger(gecho.NewConfig(gecho.WithShowCaller(true), gecho.WithLogLevel(logLevel)))
	//	mwLogger := gecho.NewLogger(gecho.NewConfig(gecho.WithShowCaller(false), gecho.WithLogLevel(logLevel)))
	logger.Info("Logger initialized")

	// Initialize database
	logger.Warn("Database logic not implemented")
	// some day

	return nil
}

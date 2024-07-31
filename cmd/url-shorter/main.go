package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod" 
)

func main() {
	cfg := config.MustLoad()
	
	log := Logger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initialize server", slog.String("address", cfg.Address))
	log.Debug("logger debug mode enabled")
}

func Logger(env string) *slog.Logger {
	var log *slog.Logger 

	switch env{
    case envLocal:
        log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envDev:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envProd:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
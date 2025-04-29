package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	go func() {
		t := time.NewTicker(500 * time.Millisecond)

		for range t.C {
			slog.Info("info message")
		}
	}()

	go func() {
		t := time.NewTicker(time.Second)

		for range t.C {
			slog.Error("error message")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
}

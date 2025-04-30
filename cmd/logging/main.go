package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("gdsahjgfaj")
	println("gadhjgajdsfg")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
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
		i := 0

		t := time.NewTicker(time.Second)

		for range t.C {
			i++
			uid := uuid.New()
			slog.Error("error message", slog.Int("count", i), slog.String("id", uid.String()))
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
}

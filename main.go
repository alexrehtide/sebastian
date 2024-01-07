package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexrehtide/sebastian/internal/application"
	configservice "github.com/alexrehtide/sebastian/internal/config-service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := application.New(configservice.New())

	if err := app.MigrateDB(); err != nil {
		fmt.Printf("error with migration: %s\r\n", err)
	}

	if err := app.Start(ctx); err != nil {
		fmt.Printf("exit reason: %s \r\n", err)
	}
}

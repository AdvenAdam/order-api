package main

import (
	"context"
	"fmt"
	"os"

	"os/signal"
	// NOTE: This depend on folder name and package name
	app "github.com/AdvenAdam/orders-api/app"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start application:", err)
	}

}

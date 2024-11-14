package main

import (
	"context"
	"fmt"

	// NOTE: This depend on folder name and package name
	app "github.com/AdvenAdam/orders-api/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start application:", err)
	}
}

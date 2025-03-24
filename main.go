package main

import (
	"context"
	"fmt"
	"go_with_redis_practice_1/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

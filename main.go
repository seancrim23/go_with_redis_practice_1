package main

import (
	"context"
	"fmt"
	"go_with_redis_practice_1/application"
)

func main() {
	app := application.New()

	err := app.Start(context.Background())
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

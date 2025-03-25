package main

import (
	"context"
	"fmt"
	"go_with_redis_practice_1/application"
	"os"
	"os/signal"
)

//TODO: all this stuff is suggested from the tutorial to expand this project
//add godotenv
//think about moving repo handler and model packages into single order package
//update repo with interface to be able to implement other data stores in future
//create a new repo using postgres
//add automated E2E testing

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

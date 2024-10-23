package main

import (
	"log"

	"github.com/dalmazox/temporal-golang/app"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer c.Close()

	w := worker.New(c, "greeting-task-queue", worker.Options{})

	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.GreetActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err.Error())
	}
}

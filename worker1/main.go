package main

import (
	"fmt"
	"log"
	"sync"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	app "github.com/quzhi1/temporal-playground"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	// Start listening to the Task Queue
	var wg sync.WaitGroup
	wg.Add(1)
	err = w.Start()
	if err != nil {
		log.Fatalln("unable to start worker 2", err)
	}
	fmt.Println("started workflow")
	wg.Wait()
}

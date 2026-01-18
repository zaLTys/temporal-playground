package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"temporal-ip-geolocation/iplocate"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Must specify a name as the command-line argument")
	}
	name := os.Args[1]

	// Match the worker configuration
	c, err := client.Dial(client.Options{
		Namespace: "default",
		HostPort:  "localhost:7234",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowID := "getAddressFromIP-" + uuid.New().String()

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: iplocate.TaskQueueName,
	}

	// Create a context with timeout for workflow execution
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	we, err := c.ExecuteWorkflow(ctx, options, iplocate.GetAddressFromIP, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	// Use the same context for getting the result
	err = we.Get(ctx, &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	fmt.Println(result)
}

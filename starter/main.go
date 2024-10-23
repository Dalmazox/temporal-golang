package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	http.HandleFunc("/wf", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        uuid.NewString() + "-greeting-wf",
		TaskQueue: "greeting-task-queue",
	}
	name := uuid.NewString()

	ctx := context.Background()

	run, err := c.ExecuteWorkflow(ctx, options, "GreetingWorkflow", name)
	if err != nil {
		log.Fatal(err.Error())
	}

	var result string
	err = run.Get(ctx, &result)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Write([]byte(result))
	w.WriteHeader(http.StatusOK)
}

type Request struct {
	Name string `json:"name"`
}

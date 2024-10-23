package app

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 100 * time.Millisecond,
	})

	ctx = workflow.WithRetryPolicy(ctx, temporal.RetryPolicy{
		MaximumAttempts: 2,
	})

	logger := workflow.GetLogger(ctx)
	var result string

	err := workflow.ExecuteActivity(ctx, GreetActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	logger.Info("workflow success")

	return result, nil
}

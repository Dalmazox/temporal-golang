package app

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func GreetActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("executing greet activity")
	return fmt.Sprintf("hello %s", name), nil
}

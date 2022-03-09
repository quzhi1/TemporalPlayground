package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result1, result2 string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result1)
	if err != nil {
		return result1, err
	}
	err = workflow.ExecuteActivity(ctx, ComposeGoodbye, name).Get(ctx, &result2)
	return result1 + "; " + result2, err
}

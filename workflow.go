package app

import (
	"fmt"
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

	future := workflow.ExecuteActivity(ctx, "ComposeGoodbye", name)
	selector := workflow.NewSelector(ctx)
	selector.AddFuture(future, func(f workflow.Future) {
		fmt.Println("selector running")
		f.Get(ctx, &result2)
	})
	selector.Select(ctx)
	return result1 + "; " + result2, err
}

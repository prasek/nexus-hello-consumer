package app

import (
	"github.com/prasek/nexus-hello-api/service"
	"go.temporal.io/sdk/workflow"
)

/*
Note: We have tentative plans to include the endpoint name in a Nexus URI, so plan to restrict this to `^[a-zA-Z][a-zA-Z0-9\-]*[a-zA-Z0-9]$` (subset of hostname RFC952) in the public preview timeframe, which will allow dashes.

Use of _ is deprecated, and will be removing support for _ in endpoint names in public preview.
*/

const (
	TaskQueue    = "my-caller-workflow-task-queue"
	endpointName = "myendpoint" // Use of _ is deprecated in endpoint names, see note above
)

func EchoCallerWorkflow(ctx workflow.Context, message string) (string, error) {
	c := workflow.NewNexusClient(endpointName, service.HelloServiceName)

	fut := c.ExecuteOperation(ctx, service.EchoOperationName, service.EchoInput{Message: message}, workflow.NexusOperationOptions{})
	var res service.EchoOutput

	var exec workflow.NexusOperationExecution
	if err := fut.GetNexusOperationExecution().Get(ctx, &exec); err != nil {
		return "", err
	}
	if err := fut.Get(ctx, &res); err != nil {
		return "", err
	}

	return res.Message, nil
}

func HelloCallerWorkflow(ctx workflow.Context, name string, language service.Language) (string, error) {
	c := workflow.NewNexusClient(endpointName, service.HelloServiceName)

	fut := c.ExecuteOperation(ctx, service.HelloOperationName, service.HelloInput{Name: name, Language: language}, workflow.NexusOperationOptions{})
	var res service.HelloOutput

	var exec workflow.NexusOperationExecution
	if err := fut.GetNexusOperationExecution().Get(ctx, &exec); err != nil {
		return "", err
	}
	if err := fut.Get(ctx, &res); err != nil {
		return "", err
	}

	return res.Message, nil
}

package app

import (
	"github.com/prasek/nexus-hello-api/service"
	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue    = "my-caller-workflow-task-queue"
	endpointName = "my_nexus_endpoint_name"
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

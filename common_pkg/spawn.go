package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type SpawnWorkflowInput struct {
	Workflow    string
	TriggerBody interface{}
}

type SpawnWorkflow struct {
}

func (SpawnWorkflow) Name() string {
	return "spawn_workflow"
}

func (SpawnWorkflow) Version() string {
	return "1.0"
}

func (SpawnWorkflow) Execute(in step.Context) (interface{}, error) {
	input := SpawnWorkflowInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	engine := in.Engine()
	err = engine.AddToQueue(input.Workflow, input.TriggerBody)
	return nil, err
}

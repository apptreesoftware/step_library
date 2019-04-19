package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
)

type FailWorkflow struct {
}

func (FailWorkflow) Name() string {
	return "fail_workflow"
}

func (FailWorkflow) Version() string {
	return "1.0"
}

func (FailWorkflow) Execute(in step.Context) (interface{}, error) {
	input := FailInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(input.Message)
}

type FailInput struct {
	Message string
}

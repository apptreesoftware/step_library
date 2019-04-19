package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
)

type QueueChildrenInput struct {
	Record    map[string]interface{}
	FieldName string
	Workflow  string
}

type QueueChildren struct {
}

func (QueueChildren) Name() string {
	return "queue_children"
}

func (QueueChildren) Version() string {
	return "1.0"
}

func (QueueChildren) Execute(ctx step.Context) (interface{}, error) {
	var input QueueChildrenInput
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return execute(input, ctx.Engine())
}

func execute(input QueueChildrenInput, engine step.Engine) (interface{}, error) {
	child := input.Record[input.FieldName]
	if arrayVal, ok := child.([]interface{}); ok && len(arrayVal) > 0 {
		for _, record := range arrayVal {
			err := engine.AddToQueue(input.Workflow, record)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
	if singleVal, ok := child.(interface{}); ok && singleVal != nil {
		err := engine.AddToQueue(input.Workflow, singleVal)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("field %s does not exist or is empty", input.FieldName))
}

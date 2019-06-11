package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type ArrayLength struct {
}

func (ArrayLength) Name() string {
	return "array_length"
}

func (ArrayLength) Version() string {
	return "1.0"
}

func (ArrayLength) Execute(ctx step.Context) (interface{}, error) {
	input := ArrayLengthInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output := ArrayLengthOutput{}
	output.Count = len(input.Array)
	return output, nil
}

type ArrayLengthInput struct {
	Array []interface{}
}

type ArrayLengthOutput struct {
	Count int
}

package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type StringLengthCounter struct {
}

func (StringLengthCounter) Name() string {
	return "string_length"
}

func (StringLengthCounter) Version() string {
	return "1.0"
}

func (StringLengthCounter) Execute(ctx step.Context) (interface{}, error) {
	input := StringLengthInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output := StringLengthOutput{}
	output.Count = len(input.Text)
	return output, nil
}

type StringLengthInput struct {
	Text string
}

type StringLengthOutput struct {
	Count int
}

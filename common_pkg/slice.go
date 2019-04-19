package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
)

type SliceString struct {
}

func (SliceString) Name() string {
	return "slice"
}

func (SliceString) Version() string {
	return "1.0"
}

func (SliceString) Execute(ctx step.Context) (interface{}, error) {
	input := SliceInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output := SliceOutput{}
	if len(input.Text) < input.EndIndex {
		return nil, errors.New("Out of bounds! The end index must be less than the text's length")
	}

	if input.StartIndex < 0 {
		return nil, errors.New("Out of bounds! The start index must be 0 or greater.")
	}

	message := input.Text[input.StartIndex:input.EndIndex]

	output.Text = message
	return output, nil
}

type SliceInput struct {
	Text       string
	StartIndex int
	EndIndex   int
}

type SliceOutput struct {
	Text string
}

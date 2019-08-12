package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"google.golang.org/api/sheets/v4"
)

type BatchWrite struct {
}

func (BatchWrite) Name() string {
	return "batch_write"
}

func (BatchWrite) Version() string {
	return "1.0"
}

func (s BatchWrite) Execute(ctx step.Context) (interface{}, error) {
	var input BatchWriteInput
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return s.execute(input)
}

func (BatchWrite) execute(input BatchWriteInput) (*BatchWriteOutput, error) {
	srv, err := ValidateInputAndGetConf(input.InputBase)
	if err != nil {
		return nil, err
	}

	sheet, err := GetSheet(input.InputBase, srv)
	if err != nil {
		return nil, err
	}

	if input.Overwrite {

	}
}

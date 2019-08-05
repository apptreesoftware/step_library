package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type BatchWrite struct {
}

func (v) Name() string {
	return "read"
}

func (BatchWrite) Version() string {
	return "1.0"
}

func (s BatchWrite) Execute(ctx step.Context) (interface{}, error) {
	// var input ReadSheetInput
	// err := ctx.BindInputs(&input)
	// if err != nil {
	// 	return nil, err
	// }
	// return s.execute(input)
}

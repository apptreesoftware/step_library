package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/robertkrimen/otto"
)

type Filter struct {
}

func (Filter) Name() string {
	return "filter"
}

func (Filter) Version() string {
	return "1.0"
}

func (Filter) Execute(ctx step.Context) (interface{}, error) {
	input := FilterInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	script := fmt.Sprintf("out = records.filter(function(record) { return %s });", input.Filter)
	vm := otto.New()

	err = vm.Set("records", input.Records)
	if err != nil {
		return nil, err
	}
	_, err = vm.Run(script)
	if err != nil {
		return nil, err
	}
	outArr, err := vm.Get("out")
	if err != nil {
		return nil, err
	}
	output := FilterOutput{}
	exportedRecords, err := outArr.Export()
	if err != nil {
		return nil, err
	}

	if outputRecords, ok := exportedRecords.([]map[string]interface{}); ok {
		output.Records = outputRecords
	}
	return output, nil
}

type FilterInput struct {
	Records []map[string]interface{}
	Filter  string
}

type FilterOutput struct {
	Records []map[string]interface{}
}

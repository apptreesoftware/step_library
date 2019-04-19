package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type SetFields struct {
}

type SetFieldsInput struct {
	Record map[string]interface{}
	Fields map[string]interface{}
}

type SetFieldsOutput struct {
	Record map[string]interface{}
}

func (SetFields) Name() string {
	return "set_fields"
}

func (SetFields) Version() string {
	return "1.0"
}

func (SetFields) Execute(ctx step.Context) (interface{}, error) {
	input := SetFieldsInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	record := input.Record
	// if no record was given we are going
	// to just write the values to an empty map
	if record == nil {
		record = make(map[string]interface{})
	}
	for k, v := range input.Fields {
		record[k] = v
	}
	return SetFieldsOutput{Record: record}, nil
}

package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
)

type ExtractFields struct {

}

func (ExtractFields) Name() string {
	return "extract_fields"
}

func (ExtractFields) Version() string {
	return "1.0"
}

func (ExtractFields) Execute(ctx step.Context) (interface{}, error) {
	inputs := ExtractFieldsInput{}
	err := ctx.BindInputs(&inputs)
	if err != nil {
		return nil, err
	}

	record := inputs.Record
	if record == nil {
		return nil, errors.New("record input is required")
	}
	if inputs.Fields == nil {
		return nil, errors.New("fields to extract is required")
	}
	subRecord := make(map[string]interface{})
	for _, fieldName := range inputs.Fields {
		subRecord[fieldName] = record[fieldName]
	}
	return ExtractFieldsOutput{Record: subRecord}, nil
}

type ExtractFieldsInput struct {
	Fields []string
	Record map[string]interface{}
}

type ExtractFieldsOutput struct {
	Record map[string]interface{}
}


package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	jsoniter "github.com/json-iterator/go"
)

type ParseJsonArray struct {
}

func (ParseJsonArray) Name() string {
	return "parse_json_array"
}

func (ParseJsonArray) Version() string {
	return "1.0"
}

func (ParseJsonArray) Execute(in step.Context) (interface{}, error) {
	input := JsonArrayInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	records := []map[string]interface{}{}
	err = jsoniter.UnmarshalFromString(input.String, &records)
	if err != nil {
		return nil, err
	}
	return JsonArrayOutput{
		Records: records,
	}, nil
}

type JsonArrayInput struct {
	String string
}

type JsonArrayOutput struct {
	Records []map[string]interface{}
}

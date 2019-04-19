package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
)

type ParseJsonObject struct {
}

func (ParseJsonObject) Name() string {
	return "parse_json_object"
}

func (ParseJsonObject) Version() string {
	return "1.0"
}

func (ParseJsonObject) Execute(ctx step.Context) (interface{}, error) {
	input := parseJsonInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	rec := map[string]interface{}{}
	err = jsoniter.UnmarshalFromString(input.String, &rec)
	if err != nil {
		return nil, err
	}
	return parseJsonOutput{Record: rec}, nil
}

type parseJsonInput struct {
	String string
}

type parseJsonOutput struct {
	Record map[string]interface{}
}

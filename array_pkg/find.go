package main

import (
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type FindInput struct {
	Records    []map[string]interface{}
	MatchField string
	MatchValue interface{}
}

type FindOutput struct {
	Record map[string]interface{}
}

type Find struct {
}

func (Find) Name() string {
	return "find"
}

func (Find) Version() string {
	return "1.0"
}

func (f Find) Execute(in step.Context) (interface{}, error) {
	input := FindInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return f.execute(input)
}

func (Find) execute(input FindInput) (*FindOutput, error) {
	for _, record := range input.Records {
		if record[input.MatchField] == input.MatchValue {
			return &FindOutput{Record: record}, nil
		}
	}
	return nil, errors.New("record not found")
}

package main

import (
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type FindInArrayInput struct {
	Records    []map[string]interface{}
	MatchField string
	MatchValue interface{}
}

type FindInArrayOutput struct {
	Record map[string]interface{}
}

type FindInArray struct {
}

func (FindInArray) Name() string {
	return "find_in_array"
}

func (FindInArray) Version() string {
	return "1.0"
}

func (f FindInArray) Execute(in step.Context) (interface{}, error) {
	input := FindInArrayInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return f.execute(input)
}

func (FindInArray) execute(input FindInArrayInput) (*FindInArrayOutput, error) {
	for _, record := range input.Records {
		if record[input.MatchField] == input.MatchValue {
			return &FindInArrayOutput{Record: record}, nil
		}
	}
	return nil, errors.New("record not found")
}

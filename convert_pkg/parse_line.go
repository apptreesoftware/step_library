package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"strconv"
	"strings"
)

type ParseLine struct {
}

func (ParseLine) Name() string {
	return "parse_line"
}

func (ParseLine) Version() string {
	return "1.0"
}

func (ParseLine) Execute(in step.Context) (interface{}, error) {
	input := splitLineInput{}

	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	components := strings.Split(input.String, input.Delimiter)
	record := map[string]interface{}{}
	for k, v := range input.StringFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		record[k] = val
	}
	for k, v := range input.IntFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		if val == "" {
			continue
		}
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse index %d (%s) into an int value", v, val)
		}
		record[k] = intVal
	}

	for k, v := range input.FloatFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		if val == "" {
			continue
		}
		floatVal, err := strconv.ParseFloat(val, 10)
		if err != nil {
			return nil, fmt.Errorf("unable to parse index %d (%s) into an int value", v, val)
		}
		record[k] = floatVal
	}

	return splitLineOutput{
		Record: record,
	}, nil
}

type splitLineInput struct {
	String       string
	Delimiter    string
	StringFields map[string]int
	IntFields    map[string]int
	FloatFields  map[string]int
}

type splitLineOutput struct {
	Record map[string]interface{}
}

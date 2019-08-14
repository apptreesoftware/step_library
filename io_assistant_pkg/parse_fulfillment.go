package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
)

type ParseFulfillmentInput struct {
	BodyAsString string `json:"bodyAsString"`
}

type ParseFulfillmentOutput struct {
	Request FulfillmentRequest `json:"request"`
}

type FulfillmentRequest struct {
	State       map[string]interface{} `json:"state"`
	UserContext map[string]interface{} `json:"userContext"`
	Context     map[string]interface{} `json:"context"`
}

type ParseFulfillment struct {
}

func (ParseFulfillment) Name() string {
	return "parse_fulfillment"
}

func (ParseFulfillment) Version() string {
	return "1.0"
}

func (p ParseFulfillment) Execute(in step.Context) (interface{}, error) {
	input := ParseFulfillmentInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return p.execute(input)
}

func (ParseFulfillment) execute(input ParseFulfillmentInput) (interface{}, error) {
	fulfillmentRequest := FulfillmentRequest{}

	err := jsoniter.Unmarshal([]byte(input.BodyAsString), &fulfillmentRequest)
	if err != nil {
		return nil, err
	}

	return ParseFulfillmentOutput{Request: fulfillmentRequest}, nil
}

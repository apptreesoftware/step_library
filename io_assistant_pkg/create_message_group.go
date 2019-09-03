package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
	"io_assistant_pkg/models"
)

type CreateMessageGroupInput struct {
	OnCompleteUrl string                   `json:"onCompleteUrl"`
	Messages      []map[string]interface{} `json:"messages"`
	UserContext   map[string]interface{}   `json:"userContext"`
	Context       map[string]interface{}   `json:"context"`
}

func (m CreateMessageGroupInput) ValidateGroup() error {
	if m.Messages == nil || len(m.Messages) == 0 {
		return errors.New("message group must contain at least 1 message")
	}
	for _, message := range m.Messages {
		if _, ok := message["id"].(string); !ok {
			return errors.New("Each message in a group must have an id")
		}
	}
	if m.OnCompleteUrl == "" {
		return errors.New("A message group must have a onCompleteUrl " +
			"that will be called when all questions in the group have been answered.")
	}
	return nil
}

type CreateMessageGroupOutput struct {
	Response models.FulfillmentResponse `json:"response"`
}

type CreateMessageGroup struct {
}

func (CreateMessageGroup) Name() string {
	return "create_message_group"
}

func (CreateMessageGroup) Version() string {
	return "1.0"
}

func (c CreateMessageGroup) Execute(in step.Context) (interface{}, error) {
	input := CreateMessageGroupInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return c.execute(input, in.Engine())
}

func (CreateMessageGroup) execute(input CreateMessageGroupInput, engine step.Engine) (*CreateMessageGroupOutput, error) {
	err := input.ValidateGroup()
	if err != nil {
		return &CreateMessageGroupOutput{}, err
	}

	response := models.NewMessageGroupResponse(
		input.Messages,
		input.OnCompleteUrl,
		input.UserContext,
		input.Context)

	return &CreateMessageGroupOutput{Response: response}, nil
}

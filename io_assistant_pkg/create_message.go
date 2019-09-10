package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io_assistant_pkg/models"
)

type CreateMessageInput struct {
	Message     map[string]interface{} `json:"message"`
	UserContext map[string]interface{} `json:"userContext"`
	Context     map[string]interface{} `json:"context"`
	Complete    bool                   `json:"complete"`
}

type CreateMessageOutput struct {
	Response models.FulfillmentResponse `json:"response"`
}

type CreateMessage struct {
}

func (CreateMessage) Name() string {
	return "create_message"
}

func (CreateMessage) Version() string {
	return "1.0"
}

func (c CreateMessage) Execute(in step.Context) (interface{}, error) {
	input := CreateMessageInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return c.execute(input, in.Engine())
}

func (CreateMessage) execute(input CreateMessageInput, engine step.Engine) (*CreateMessageOutput, error) {
	response := models.NewMessageResponse(input.Message, input.UserContext, input.Context, input.Complete)
	return &CreateMessageOutput{Response: response}, nil
}

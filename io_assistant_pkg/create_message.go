package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io_assistant_pkg/models"
)

type CreateMessageInput struct {
	Message     models.MessageInput    `json:"message"`
	UserContext map[string]interface{} `json:"userContext"`
	Context     map[string]interface{} `json:"context"`
}

type CreateMessageOutput struct {
	Response Response
}

type Response struct {
	Message     models.MessageData     `json:"message"`
	UserContext map[string]interface{} `json:"userContext"`
	Context     map[string]interface{} `json:"context"`
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
	err := input.Message.ValidateMessageInput()
	if err != nil {
		return &CreateMessageOutput{}, err
	}

	workflowUrl := ""
	if input.Message.OnCompleteWorkflow != "" {
		workflowUrl, err = engine.GetWorkflowUrl(input.Message.OnCompleteWorkflow, nil)
		if err != nil {
			return &CreateMessageOutput{}, err
		}
	}

	response := Response{
		Message: models.MessageData{
			MessageBase: input.Message.MessageBase,
			OnCompleteUrl: workflowUrl,
		},
		UserContext: input.UserContext,
		Context: input.Context,
	}

	return &CreateMessageOutput{Response: response}, nil
}

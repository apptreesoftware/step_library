package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io_assistant_pkg/models"
)

type CreateMessageInput struct {
	Message            models.MessageBase     `json:"message"`
	UserContext        map[string]interface{} `json:"userContext"`
	Context            map[string]interface{} `json:"context"`
	OnCompleteWorkflow string                 `json:"onCompleteWorkflow"`
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
	err := input.Message.ValidateMessageInput()
	if err != nil {
		return &CreateMessageOutput{}, err
	}

	workflowUrl := ""
	if input.OnCompleteWorkflow != "" {
		workflowUrl, err = engine.GetWorkflowUrl(input.OnCompleteWorkflow, nil)
		if err != nil {
			return &CreateMessageOutput{}, err
		}
	}
	response := models.NewMessageResponse(input.Message, workflowUrl, input.UserContext, input.Context)
	return &CreateMessageOutput{Response: response}, nil
}

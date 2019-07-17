package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
	"io_assistant_pkg/models"
)

type CreateMessageGroupInput struct {
	OnCompleteWorkflow string                 `json:"onCompleteWorkflow"`
	Messages           []models.MessageOption `json:"messages"`
	UserContext        map[string]interface{} `json:"userContext"`
	Context            map[string]interface{} `json:"context"`
}

func (m CreateMessageGroupInput) ValidateGroup() error {
	if m.Messages == nil || len(m.Messages) == 0 {
		return errors.New("message group must contain at least 1 message")
	}
	for _, message := range m.Messages {
		err := message.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

type CreateMessageGroupOutput struct {
	Response models.FulfillmentResponse
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

	workflowUrl := ""
	if input.OnCompleteWorkflow != "" {
		workflowUrl, err = engine.GetWorkflowUrl(input.OnCompleteWorkflow, nil)
		if err != nil {
			return &CreateMessageGroupOutput{}, err
		}
	}

	response := models.NewMessageGroupResponse(input.Messages, workflowUrl, input.UserContext, input.Context)

	return &CreateMessageGroupOutput{Response: response}, nil
}

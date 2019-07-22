package main

import (
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io_assistant_pkg/models"
)

type BuildOptionsInput struct {
	Records   []map[string]interface{} `json:"records"`
	IdField   string                   `json:"idField"`
	TextField string                   `json:"textField"`
}

type BuildOptionsOutput struct {
	Options []models.MessageOption `json:"options"`
}

type BuildOptions struct {
}

func (BuildOptions) Name() string {
	return "build_select_options"
}

func (BuildOptions) Version() string {
	return "1.0"
}

func (b BuildOptions) Execute(in step.Context) (interface{}, error) {
	input := BuildOptionsInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return b.execute(input)
}

func (BuildOptions) execute(input BuildOptionsInput) (*BuildOptionsOutput, error) {
	options := make([]models.MessageOption, 0)
	if input.Records == nil || len(input.Records) == 0 {
		return &BuildOptionsOutput{Options: options}, nil
	}

	for _, record := range input.Records {
		id := ""
		if idString, ok := record[input.IdField].(string); ok {
			id = idString
		}

		text := ""
		if textString, ok := record[input.TextField].(string); ok {
			text = textString
		}

		if id == "" || text == "" {
			return nil, errors.New("records are missing ID or text field or those fields are not strings")
		}

		options = append(options, models.MessageOption{Id: id, Text: text})
	}
	return &BuildOptionsOutput{Options: options}, nil
}

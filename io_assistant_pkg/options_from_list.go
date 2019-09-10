package main

import (
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io_assistant_pkg/models"
)

type OptionsListInput struct {
	Records []string `json:"records"`
}

type BuildOptionsList struct {
}

func (BuildOptionsList) Name() string {
	return "build_options_list"
}

func (BuildOptionsList) Version() string {
	return "1.0"
}

func (b BuildOptionsList) Execute(in step.Context) (interface{}, error) {
	input := OptionsListInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return b.execute(input)
}

func (BuildOptionsList) execute(input OptionsListInput) (*BuildOptionsOutput, error) {
	if input.Records == nil || len(input.Records) == 0 {
		return nil, errors.New("no records provided to build options from")
	}
	options := make([]models.MessageOption, len(input.Records))
	for idx, record := range input.Records {
		options[idx] = models.MessageOption{
			Id:   record,
			Text: record,
		}
	}

	return &BuildOptionsOutput{Options: options}, nil
}

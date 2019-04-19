package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
)

type WriteFile struct {}

func (WriteFile) Name() string {
	return "write_file"
}

func (WriteFile) Version() string {
	return "1.0"
}

func (WriteFile) Execute(in step.Context) (interface{}, error) {
	input := &WriteFileInput{}
	err := in.BindInputs(input)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(input.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file.WriteString(input.Contents)
	file.WriteString("\n")

	return WriteFileOutput{}, nil
}

type WriteFileInput struct {
	FilePath string
	Contents string
}

type WriteFileOutput struct {}

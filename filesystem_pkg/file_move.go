package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type FileMoveInput struct {
	FilePath    string
	ToDirectory string
	FileName    string
}

type FileMoveOutput struct {
	Success bool
}

type FileMove struct {

}

func (FileMove) Name() string {
	return "file_move"
}

func (FileMove) Version() string {
	return "1.0"
}

func (f FileMove) Execute(in step.Context) (interface{}, error) {
	var input FileMoveInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return f.execute(input)
}

func (FileMove) execute(input FileMoveInput) (FileMoveOutput, error) {
	if input.FilePath == "" {
		return FileMoveOutput{Success:false}, errors.New("file can not be blank")
	}
	if input.ToDirectory == "" {
		return FileMoveOutput{Success:false}, errors.New("new directory can not be blank")
	}
	if input.FileName == "" {
		_, fileName := filepath.Split(input.FilePath)
		input.FileName = fileName
	}
	newPath := filepath.Join(input.ToDirectory, input.FileName)
	err := os.Rename(input.FilePath, newPath)
	if err != nil {
		return FileMoveOutput{}, err
	}
	return FileMoveOutput{Success: true}, nil
}




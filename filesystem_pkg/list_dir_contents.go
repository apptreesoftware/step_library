package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"path/filepath"
	"regexp"
)

type ListDirectory struct{}

func (ListDirectory) Name() string {
	return "list_directory_contents"
}

func (ListDirectory) Version() string {
	return "1.0"
}

func (f ListDirectory) Execute(ctx step.Context) (interface{}, error) {
	input := ListDirectoryInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return f.execute(input)
}

func (ListDirectory) execute(input ListDirectoryInput) (ListDirectoryOutput, error) {
	dir, err := os.Open(input.DirectoryPath)
	if err != nil {
		return ListDirectoryOutput{}, err
	}
	files, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		return ListDirectoryOutput{}, err
	}

	var reg *regexp.Regexp
	if len(input.MatchPattern) > 0 {
		regMatch, err := regexp.Compile(input.MatchPattern)
		if err != nil {
			return ListDirectoryOutput{}, fmt.Errorf("Invalid MatchPattern %s", input.MatchPattern)
		}
		reg = regMatch
	}

	var output []string
	for _, file := range files {
		_, fileName := filepath.Split(file.Name())
		if reg != nil && !reg.MatchString(fileName) {
			continue
		}

		if !input.Relative {
			output = append(output, filepath.Join(input.DirectoryPath, file.Name()))
		} else {
			output = append(output, file.Name())
		}
	}

	return ListDirectoryOutput{
		Files: output,
	}, nil
}

type ListDirectoryInput struct {
	DirectoryPath string
	MatchPattern  string
	Relative      bool
}

type ListDirectoryOutput struct {
	Files []string
}

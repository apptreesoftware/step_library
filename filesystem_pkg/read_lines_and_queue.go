package main

import (
	"bufio"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"path/filepath"
)

type ReadLinesAndQueue struct{}

func (ReadLinesAndQueue) Name() string {
	return "read_lines_and_queue"
}

func (ReadLinesAndQueue) Version() string {
	return "1.0"
}

func (f ReadLinesAndQueue) Execute(ctx step.Context) (interface{}, error) {
	input := ReadLinesAndQueueInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	err = f.execute(input, ctx)
	return nil, err
}

func (ReadLinesAndQueue) execute(input ReadLinesAndQueueInput, ctx step.Context) error {
	moveTo := input.MoveToDirectoryAfterProcessing
	for _, path := range input.FilePaths {
		if ctx.Environment().Debug {
			println("reading file", path)
		}
		file, err := os.Open(path)
		if err != nil {
			_ = file.Close()
			return err
		}

		scanner := bufio.NewScanner(file)
		engine := ctx.Engine()

		for scanner.Scan() {
			lineRecord := scanner.Text()
			err := engine.AddToQueue(input.Workflow, lineRecord)
			if err != nil {
				_ = file.Close()
				return err
			}
			if ctx.Environment().Debug {
				println("Queuing line", lineRecord)
			}
		}
		_ = file.Close()
		if len(moveTo) > 0 {
			if _, err := os.Stat(moveTo); err != nil {
				_ = os.MkdirAll(moveTo, 0777)
			}
			_, name := filepath.Split(path)
			movePath := filepath.Join(moveTo, name)
			err := os.Rename(path, filepath.Join(moveTo, name))
			if err != nil {
				return fmt.Errorf("Unable to move file to %s", movePath)
			}
		}
	}
	return nil
}

type ReadLinesAndQueueInput struct {
	FilePaths                      []string
	Workflow                       string
	MoveToDirectoryAfterProcessing string
}

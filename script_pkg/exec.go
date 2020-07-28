package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"

	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type ExecRunner struct {
}

func (ExecRunner) Name() string {
	return "exec"
}

func (ExecRunner) Version() string {
	return "1.0"
}

func (j ExecRunner) Execute(ctx step.Context) (interface{}, error) {
	input := ExecInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, errors.Errorf("Unable to bind inputs: %v", err)
	}

	return j.execute(input, ctx.Environment().Debug)
}

func (j ExecRunner) execute(input ExecInput, debug bool) (ExecOutput, error) {
	execCmd := exec.Command(input.Cmd, input.Args...)
	if debug {
		println(fmt.Sprintf("Executing CMD: %s", execCmd.String()))
	}
	if output, err := execCmd.Output(); err == nil {
		outputStr := string(output)
		if debug {
			println(fmt.Sprintf("Output: %s", outputStr))
		}
		return ExecOutput{
			Output: outputStr,
		}, nil
	} else {
		return ExecOutput{}, errors.Errorf("Unable to execute command: %v", err)
	}
}

type ExecInput struct {
	Cmd string
	Args []string
	Dir string
}

type ExecOutput struct {
	Output string
}

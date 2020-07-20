package main

import "testing"

func TestExec(t *testing.T) {
	input := ExecInput{
		Cmd:  "apptree",
		Args: []string{"get", "workflows", "--project", "uark"},
	}

	runner := ExecRunner{}
	output, err := runner.execute(input, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	val := output.Output
	println(val)
}

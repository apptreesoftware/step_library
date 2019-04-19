package db_common

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type MockExecute struct {
}

func (MockExecute) Name() string {
	return "mock_execute"
}

func (MockExecute) Version() string {
	return "1.0"
}

func (MockExecute) Execute(in step.Context) (interface{}, error) {
	cmd := DatabaseCommand{}
	err := in.BindInputs(&cmd)
	if err != nil {
		return nil, err
	}

	println("CONNECTION:", cmd.ConnectionString)
	println("MOCK:", cmd.Sql)

	return nil, nil
}

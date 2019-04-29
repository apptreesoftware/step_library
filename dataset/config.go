package dataset

import "github.com/apptreesoftware/go-workflow/pkg/step"

type Config struct {
}

func (Config) Name() string {
	return "describe"
}

func (Config) Version() string {
	return "1.0"
}

func (Config) Execute(in step.Context) (interface{}, error) {
	input := ConfigInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"time"
)

type Sleep struct {
}

func (Sleep) Name() string {
	return "sleep"
}

func (Sleep) Version() string {
	return "1.0"
}

func (Sleep) Execute(in step.Context) (interface{}, error) {
	input := sleepInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	if input.Seconds != 0 {
		time.Sleep(time.Duration(input.Seconds) * time.Second)
	} else if input.Millis != 0 {
		time.Sleep(time.Duration(input.Millis) * time.Millisecond)
	}
	return nil, nil
}

type sleepInput struct {
	Seconds int64
	Millis  int64
}

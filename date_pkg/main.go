package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"time"
)

func main() {
	step.Register(EpocDateParser{})
	step.Run()
}

type EpocDateParser struct {
}

func (e EpocDateParser) Name() string {
	return "from_epoch"
}

func (e EpocDateParser) Version() string {
	return "1.0"
}

func (e EpocDateParser) Execute(in step.Context) (interface{}, error) {
	var inputs epocDateInput
	err := in.BindInputs(&inputs)
	if err != nil {
		return nil, err
	}
	var date time.Time

	if inputs.Seconds != 0 {
		date = time.Unix(inputs.Seconds, 0)
	} else if inputs.Millis != 0 {
		date = time.Unix(0, inputs.Millis*1000000)
	} else if inputs.Nanos != 0 {
		date = time.Unix(0, inputs.Nanos)
	}
	return epochDateOutput{
		Date: date,
	}, nil
}

type epocDateInput struct {
	Seconds int64
	Millis  int64
	Nanos   int64
}

type epochDateOutput struct {
	Date time.Time
}

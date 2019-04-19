package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(Filter{})
	step.Register(StringLengthCounter{})
	step.Register(SliceString{})
	step.Register(ObjectCompare{})
	step.Register(FailWorkflow{})
	step.Register(Sleep{})
	step.Run()
}

package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(BuildOptions{})
	step.Register(CreateMessageGroup{})
	step.Register(CreateMessage{})
	step.Register(ParseFulfillment{})

	step.Run()
}

package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(CreateMessage{})
	step.Register(ParseFulfillment{})

	step.Run()
}

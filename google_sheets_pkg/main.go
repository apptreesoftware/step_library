package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(ReadSheet{})
	step.Register(WriteToSheet{})
	step.Run()
}

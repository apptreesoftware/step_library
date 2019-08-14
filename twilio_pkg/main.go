package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

func main() {
	step.Register(SendSms{})
	step.Register(SendSMSCopilot{})
	step.Run()
}

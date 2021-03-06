package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/sfreiberg/gotwilio"
	"golang.org/x/xerrors"
)

type SendSmsBase struct {
	TwilioAccountSid string
	TwilioAuthToken  string
	Message          string
	To               string
}

type SendSmsInput struct {
	SendSmsBase
	From string
}

type SendSmsOutput struct {
	Success      bool
	MessageId    string
	ErrorMessage string
}

type SendSms struct {
}

func (SendSms) Name() string {
	return "send_sms"
}

func (SendSms) Version() string {
	return "1.0"
}

func (s SendSms) Execute(in step.Context) (interface{}, error) {
	input := SendSmsInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return s.execute(input)
}

func (s SendSms) execute(input SendSmsInput) (interface{}, error) {
	twilio := gotwilio.NewTwilioClient(input.TwilioAccountSid, input.TwilioAuthToken)

	res, exception, err := twilio.SendSMS(input.From, input.To, input.Message, "", "")
	if exception != nil {
		return &SendSmsOutput{Success: false, ErrorMessage: exception.MoreInfo}, err
	}
	if err != nil {
		return &SendSmsOutput{Success: false, ErrorMessage: err.Error()}, err
	}
	return &SendSmsOutput{Success: true, MessageId: res.Sid}, nil
}

type SendSmsCopilotInput struct {
	SendSmsBase
	MessageServiceSid string
}

type SendSMSCopilot struct {
}

func (SendSMSCopilot) Name() string {
	return "send_sms_copilot"
}

func (SendSMSCopilot) Version() string {
	return "1.0"
}

func (s SendSMSCopilot) Execute(in step.Context) (interface{}, error) {
	var input SendSmsCopilotInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, xerrors.Errorf("Unable to read inputs: %v", err)
	}

	return s.execute(input)
}

func (SendSMSCopilot) execute(input SendSmsCopilotInput) (interface{}, error) {
	twilio := gotwilio.NewTwilioClient(input.TwilioAccountSid, input.TwilioAuthToken)

	res, exception, err := twilio.SendSMSWithCopilot(input.MessageServiceSid, input.To, input.Message, "", "")
	if exception != nil {
		return &SendSmsOutput{Success: false, ErrorMessage: exception.MoreInfo}, err
	}
	if err != nil {
		return &SendSmsOutput{Success: false, ErrorMessage: err.Error()}, err
	}
	return &SendSmsOutput{Success: true, MessageId: res.Sid}, nil
}

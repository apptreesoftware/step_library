package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/sfreiberg/gotwilio"
)

type SendSmsInput struct {
	TwilioAccountSid  string
	TwilioAuthToken   string
	Message           string
	From              string
	To                string
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

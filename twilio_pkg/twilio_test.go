package main

import (
	"errors"
	"os"
	"testing"
)

func TestSendSms_Execute(t *testing.T) {
	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM")
	to := os.Getenv("TWILIO_TO_TEST")

	input := SendSmsInput{
		SendSmsBase: SendSmsBase{
			TwilioAuthToken:  twilioAuthToken,
			TwilioAccountSid: twilioAccountSid,
			To:               to,
			Message:          "This is a test message from unit tests. Please disregard.",
		},
		From: from,
	}

	sendsms := SendSms{}
	response, err := sendsms.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if smsResponse, ok := response.(*SendSmsOutput); ok {
		if !smsResponse.Success {
			t.Error(errors.New(smsResponse.ErrorMessage))
			t.FailNow()
		}
		return
	}
	t.Error(errors.New("output is not SendSmsOutput"))
	t.Fail()
}

func TestSendSmsCopilot_Execute(t *testing.T) {
	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	to := os.Getenv("TWILIO_TO_TEST")
	serviceSid := os.Getenv("COPILOT_SID")

	input := SendSmsCopilotInput{
		SendSmsBase: SendSmsBase{
			TwilioAuthToken:  twilioAuthToken,
			TwilioAccountSid: twilioAccountSid,
			To:               to,
			Message:          "This is a test message from unit tests. Please disregard.",
		},
		MessageServiceSid: serviceSid,
	}

	sendsms := SendSMSCopilot{}
	response, err := sendsms.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if smsResponse, ok := response.(*SendSmsOutput); ok {
		if !smsResponse.Success {
			t.Error(errors.New(smsResponse.ErrorMessage))
			t.FailNow()
		}
		return
	}
	t.Error(errors.New("output is not SendSmsOutput"))
	t.Fail()
}

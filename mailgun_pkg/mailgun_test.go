package main

import (
	"errors"
	"os"
	"testing"
)

func TestSendEmail_Execute(t *testing.T) {
	domain := os.Getenv("MAILGUN_DOMAIN")
	username := os.Getenv("MAILGUN_USERNAME")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	subject := "Unit Testing Message"
	ptMessage := "Hello User,\n\nThis is a message from unit tests. Please disregard and know emails are functioning.\n\nAppTree Unit Tests"
	htmlMessage := "<h1>Hello User</h1><br/><br/><p>This is a message from unit tests. Please disregard and know emails are functioning.</p><br/><br/>AppTree Unit Tests"
	to := os.Getenv("MAILGUN_TO_TEST")

	input := SendEmailInput{
		Domain: domain,
		SenderUsername: username,
		ApiKey: apiKey,
		Subject: subject,
		PlainTextMessage: ptMessage,
		HtmlMessage: htmlMessage,
		To: to,
	}

	sendEmail := SendEmail{}
	resp, err := sendEmail.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if emailResp, ok := resp.(*SendEmailOutput); ok {
		if !emailResp.Success {
			t.Error(errors.New(emailResp.ErrorMessage))
			t.FailNow()
		}
		return
	}
	t.Error(errors.New("response is not SendEmailResponse"))
	t.Fail()
}

package main

import (
	"context"
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

type SendEmailInput struct {
	Domain           string
	SenderUsername   string
	ApiKey           string
	Subject          string
	HtmlMessage      string
	PlainTextMessage string
	To               string
}

type SendEmailOutput struct {
	Success      bool
	EmailId      string
	ErrorMessage string
}

type SendEmail struct {
}

func (SendEmail) Name() string {
	return "send_email"
}

func (SendEmail) Version() string {
	return "1.0"
}

func (s SendEmail) Execute(in step.Context) (interface{}, error) {
	input := SendEmailInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return s.execute(input)
}

func (s SendEmail) execute(input SendEmailInput) (interface{}, error) {
	if input.PlainTextMessage == "" {
		return nil, errors.New("plain text message is required")
	}

	if input.To == "" {
		return nil, errors.New("to parameter is required")
	}

	mg := mailgun.NewMailgun(input.Domain, input.ApiKey)
	message := mg.NewMessage(input.SenderUsername, input.Subject, input.PlainTextMessage, input.To)
	if input.HtmlMessage != "" {
		message.SetHtml(input.HtmlMessage)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	if err != nil {
		return &SendEmailOutput{Success: false, ErrorMessage: err.Error()}, err
	}
	return &SendEmailOutput{Success: true, EmailId: id}, nil
}



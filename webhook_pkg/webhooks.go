package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostWebhookInput struct {
	Url         string
	Body        string
	ContentType string
	Header      http.Header //map[string][]string
}

type GetWebhookInput struct {
	Url     string
	Header http.Header
}

type WebhookOutput struct {
	ResponseBody string
}

type PostWebhook struct {
}

func (PostWebhook) Name() string {
	return "post"
}

func (PostWebhook) Description() string {
	return "Posts a webhook"
}

func (PostWebhook) Version() string {
	return "1.0"
}

func (PostWebhook) Execute(ctx step.Context) (interface{}, error) {
	input := PostWebhookInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(input.Body)
	request, err := http.NewRequest(http.MethodPost, input.Url, reader)
	if err != nil {
		return nil, err
	}
	request.Header = input.Header
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return handleResponse(resp, err)
}

type GetWebhook struct {
}

func (GetWebhook) Name() string {
	return "get"
}

func (GetWebhook) Description() string {
	return "Performs a GET webhook"
}

func (GetWebhook) Version() string {
	return "1.0"
}

func (GetWebhook) Execute(ctx step.Context) (interface{}, error) {
	input := GetWebhookInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, input.Url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = input.Header
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return handleResponse(resp, err)
}

func handleResponse(resp *http.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return &WebhookOutput{ResponseBody: string(body)}, nil
	}
	return nil, fmt.Errorf("invalid response code %d", resp.StatusCode)
}

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
	"net/http"
	"net/url"
)

type ShortenUrlInput struct {
	FirebaseApiKey    string
	FirebaseUrlPrefix string
	Url               string
}

type ShortenUrlOutput struct {
	ShortUrl string
}

type UrlShortenRequest struct {
	LongDynamicLink string `json:"longDynamicLink"`
}

type UrlShortenResponse struct {
	ShortLink string        `json:"shortLink"`
	Error     *FirebaseError `json:"error"`
}

type FirebaseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ShortenUrl struct {
}

func (ShortenUrl) Name() string {
	return "shorten_url"
}

func (ShortenUrl) Version() string {
	return "1.0"
}

func (s ShortenUrl) Execute(in step.Context) (interface{}, error) {
	input := ShortenUrlInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return s.execute(input)
}

func (s ShortenUrl) execute(input ShortenUrlInput) (interface{}, error) {
	escapedUrl := url.QueryEscape(input.Url)
	requestBody := UrlShortenRequest{
		LongDynamicLink: fmt.Sprintf("%s/?link=%s", input.FirebaseUrlPrefix, escapedUrl),
	}
	jsonBody, err := jsoniter.Marshal(&requestBody)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	uri, err := url.Parse("https://firebasedynamiclinks.googleapis.com/v1/shortLinks")
	if err != nil {
		return nil, err
	}
	uri.RawQuery = fmt.Sprintf("key=%s", input.FirebaseApiKey)

	resp, err := client.Post(uri.String(), "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result UrlShortenResponse
	err = jsoniter.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, errors.New(result.Error.Message)
	}

	return &ShortenUrlOutput{ShortUrl: result.ShortLink}, nil
}

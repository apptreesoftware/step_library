package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Tweet struct{}

func (Tweet) Name() string {
	return "tweet"
}

func (Tweet) Version() string {
	return "1.0"
}

func (Tweet) Execute(ctx step.Context) (interface{}, error) {
	input := TweetInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	output := TweetOutput{}

	credentials := Credentials{
		AccessToken:       input.AccessToken,
		AccessTokenSecret: input.AccessTokenSecret,
		ConsumerKey:       input.ConsumerKey,
		ConsumerSecret:    input.ConsumerSecret,
	}

	client, err := GetUserClient(&credentials)
	if err != nil {
		return nil, err
	}

	_, _, err = client.Statuses.Update(
		input.Text, nil)
	if err != nil {
		return nil, err
	}

	output.Success = true
	return output, nil
}

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetUserClient(credentials *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(credentials.ConsumerKey,
		credentials.ConsumerSecret)
	token := oauth1.NewToken(credentials.AccessToken, credentials.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type TweetInput struct {
	Text              string
	AccessToken       string
	AccessTokenSecret string
	ConsumerKey       string
	ConsumerSecret    string
}

type TweetOutput struct {
	Success bool
}

package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
	"net/http"
)

type GetCount struct {
	Fetch
}

func (g GetCount) Name() string {
	return "get_count"
}

func (g GetCount) Version() string {
	return "1.0"
}

func (g GetCount) Execute(in step.Context) (interface{}, error) {
	input := FetchInput{}
	err := in.BindInputs(&input)
	input.Count = true

	if in.Environment().Debug {
		println("URL", input.Url)
	}

	if err != nil {
		return nil, err
	}
	auth, err := g.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}
	g.authItem = auth
	if err != nil {
		return nil, err
	}
	uri, err := g.buildUrl(input)
	if err != nil {
		return nil, err
	}

	req, err := g.buildRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, xerrors.Errorf("The count request returned a status of %d", resp.StatusCode)
	}
	dec := jsoniter.NewDecoder(resp.Body)
	jsonMap := JsonMap{}
	err = dec.Decode(&jsonMap)
	if err != nil {
		return nil, xerrors.Errorf("Unable to decode response: %v", err)
	}
	if count, ok := jsonMap["@odata.count"].(float64); ok {
		return CountOutput{Count: int(count)}, nil
	}
	return nil, xerrors.Errorf("Unable to parse count from response")
}

type CountOutput struct {
	Count int
}

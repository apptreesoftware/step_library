package main

import (
	"bytes"
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

const createRequestMethod = "POST"

type CreateRecord struct {
	Fetch
}

func (CreateRecord) Name() string {
	return "create_record"
}

func (CreateRecord) Version() string {
	return "1.0"
}

func (create CreateRecord) ExecuteJson(jsonString string) (interface{}, error) {
	input := &Facility360CreateIn{}
	err := json.Unmarshal([]byte(jsonString), input)
	if err != nil {
		return nil, err
	}
	return create.execute(input)
}

func (create CreateRecord) Execute(in step.Context) (interface{}, error) {
	input := &Facility360CreateIn{}
	err := in.BindInputs(input)
	if err != nil {
		return nil, err
	}
	return create.execute(input)
}

func (create CreateRecord) execute(input *Facility360CreateIn) (interface{}, error) {
	// get authenticated
	err := create.LogMeInFacility360(input.Facility360Input)
	if err != nil {
		return nil, err
	}

	createUrl, err := create.getUrl(input.Url, input.Endpoint)
	if err != nil {
		return nil, err
	}

	data, err := create.getRecordData(input)
	if err != nil {
		return nil, err
	}

	req, err := create.buildRequest(createRequestMethod, createUrl.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	resp, err := create.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		// if status code is not 200
		// read body as string and return
		return create.handleFailedResponse(resp)
	}
	defer resp.Body.Close()
	return create.handleUpsertResponse(resp)
}

func (create CreateRecord) getRecordData(input *Facility360CreateIn) ([]byte, error) {
	return json.Marshal(input.Record)
}

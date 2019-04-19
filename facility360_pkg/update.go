package main

import (
	"bytes"
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"net/url"
	"strconv"
)

const key = "key"
const updateRequestMethod = "PATCH"

type UpdateRecord struct {
	Fetch
}

func (UpdateRecord) Name() string {
	return "update_record"
}

func (UpdateRecord) Version() string {
	return "1.0"
}

func (update UpdateRecord) ExecuteJson(jsonString string) (interface{}, error) {
	input := &Facility360UpdateIn{}
	err := json.Unmarshal([]byte(jsonString), input)
	if err != nil {
		return nil, err
	}
	return update.execute(input)
}

func (update UpdateRecord) Execute(in step.Context) (interface{}, error) {
	input := &Facility360UpdateIn{}
	err := in.BindInputs(input)
	if err != nil {
		return nil, err
	}
	return update.execute(input)
}

func (update UpdateRecord) execute(input *Facility360UpdateIn) (interface{}, error) {
	err := update.LogMeInFacility360(input.Facility360Input)
	if err != nil {
		return nil, err
	}
	updateUrl, err := update.getUrl(input.Url, input.Endpoint)
	// add update key to url
	updateUrl = update.addKeyToUrl(updateUrl, input.Id)
	if err != nil {
		return nil, err
	}
	data, err := update.getRecordData(input)
	if err != nil {
		return nil, err
	}
	req, err := update.buildRequest(updateRequestMethod, updateUrl.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	resp, err := update.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return update.handleFailedResponse(resp)
	}
	defer resp.Body.Close()
	return update.handleUpsertResponse(resp)
}

func (update UpdateRecord) getRecordData(input *Facility360UpdateIn) ([]byte, error) {
	return json.Marshal(input.Record)
}

func (update UpdateRecord) addKeyToUrl(url *url.URL, id int) *url.URL {
	query := url.Query()
	query.Add(key, strconv.Itoa(id))
	url.RawQuery = query.Encode()
	return url
}

package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"testing"
)

func Test_FetchPaged(t *testing.T) {
	fetchInput := FetchInput{
		Username: os.Getenv("famis_user"),
		Password: os.Getenv("famis_pass"),
		Url:      os.Getenv("famis_url"),
		Endpoint: "MobileWebServices/apis/360facility/v1/spaces",
	}

	fetcher := Fetch{}
	out, err := fetcher.execute(fetchInput)
	if err != nil {
		t.Error(err.Error())
		return
	}

	output, ok := out.(FetchListOutputs)
	if !ok {
		t.Log("Output is not type FetchListOutputs")
		t.Fail()
	}
	if len(output.Records) != output.Count {
		t.Log("Output count does not match record count")
		t.Fail()
	}

	if output.Count == 0 {
		t.Log("No records were returned")
		t.Fail()
	}
	println("Records returned:", output.Count)
}

func Test_WOPaged(t *testing.T) {
	fetchInput := FetchInput{
		Username: os.Getenv("famis_user"),
		Password: os.Getenv("famis_pass"),
		Url:      os.Getenv("famis_url"),
		Endpoint: "MobileWebServices/apis/360facility/v1/workorders",
		Filter:   "StatusId eq 1",
		Select:   "StatusId,Id",
	}

	fetcher := Fetch{}
	out, err := fetcher.execute(fetchInput)
	if err != nil {
		t.Error(err.Error())
		return
	}

	output, ok := out.(FetchListOutputs)
	if !ok {
		t.Log("Output is not type FetchListOutputs")
		t.Fail()
	}
	if len(output.Records) != output.Count {
		t.Log("Output count does not match record count")
		t.Fail()
	}

	if output.Count == 0 {
		t.Log("No records were returned")
		t.Fail()
	}
	println("Records returned:", output.Count)
}

func Test_SingleFetch(t *testing.T) {
	fetchInput := FetchInput{
		Username: os.Getenv("famis_user"),
		Password: os.Getenv("famis_pass"),
		Url:      os.Getenv("famis_url"),
		Endpoint: "MobileWebServices/apis/360facility/v1/workorders",
		Filter:   "Id eq 3",
	}

	fetcher := GetRecord{}
	out, err := fetcher.execute(fetchInput, true)
	if err != nil {
		t.Error(err.Error())
		return
	}

	output, ok := out.(FetchSingleOutput)
	if !ok {
		t.Log("Output is not type FetchListOutputs")
		t.Fail()
	}
	if output.Found == false {
		t.Log("Found == false")
		t.Fail()
	}
	if output.Record["Id"] != float64(3) {
		t.Log("Id does not == 3")
		t.Fail()
	}
	fmt.Printf("%+v\n", output.Record)
}

func Test_GetRecordsAndQueue(t *testing.T) {
	fetchInput := FetchAndQueueInput{
		FetchInput: FetchInput{
			Username: os.Getenv("famis_user"),
			Password: os.Getenv("famis_pass"),
			Url:      os.Getenv("famis_url"),
			Endpoint: "MobileWebServices/apis/360facility/v1/workorders",
			Filter:   "StatusId eq 1",
		},
		Workflow: "Test",
	}

	fetcher := FetchAndQueue{}
	_, err := fetcher.executeQueue(fetchInput, step.IpcContext{})
	if err != nil {
		t.Error(err.Error())
		return
	}

}

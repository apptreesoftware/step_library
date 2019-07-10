package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/core"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"testing"
)

func TestShortenUrl_Execute(t *testing.T) {
	firebaseApi := os.Getenv("FIREBASE_API_KEY")
	urlPrefix := os.Getenv("FIREBASE_URL_PREFIX")
	origUrl := "https://google.com/?someKey=someVal&thisKey=thisVal"

	input := ShortenUrlInput{
		FirebaseApiKey:    firebaseApi,
		FirebaseUrlPrefix: urlPrefix,
		Url:               origUrl,
	}

	shorten := ShortenUrl{}
	resp, err := shorten.execute(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output, ok := resp.(*ShortenUrlOutput); ok {
		if output.ShortUrl == "" {
			t.Error(errors.New("url is empty"))
			t.FailNow()
		}
		return
	}
	t.Error("response was not s ShortenUrlResponse")
	t.Fail()
}

func TestQuery(t *testing.T) {
	input := QueryInput{
		QueryParameters: getQueryParams(),
	}
	input.ServiceAccountJson = os.Getenv("FIREBASE_JSON")
	input.CollectionPath = "entries"

	step := Query{}
	res, err := step.execute(input, "runString")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	println(fmt.Sprintf("%d records found", len(res.Records)))
	for _, record := range res.Records {
		myString, err := json.Marshal(&record)
		if err == nil {
			println(string(myString))
		}
	}
}

func TestQueryAndQueue(t *testing.T) {
	engine := createEngine()
	input := QueryQueueInput{
		QueryParameters: getQueryParams(),
		Workflow:        "fake_workflow_two",
	}
	input.ServiceAccountJson = os.Getenv("FIREBASE_JSON")
	input.CollectionPath = "apps"

	step := QueryAndQueue{}
	res, err := step.execute(input, "runString", engine)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	println(fmt.Sprintf("%d records found", len(res.Records)))
	for _, record := range res.Records {
		myString, err := json.Marshal(&record)
		if err == nil {
			println(string(myString))
		}
	}
}

func TestUpdate(t *testing.T) {
	input := UpdateDocumentInput{
		QueryParameters: getQueryParams(),
		UpdateFields:    map[string]interface{}{"newField": "newStringVal"},
		MaxUpdateSize:   1,
	}
	input.ServiceAccountJson = os.Getenv("FIREBASE_JSON")
	input.CollectionPath = "apps"

	step := UpdateDocument{}
	res, err := step.execute(input, "runString")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if res.RecordsUpdated != 1 {
		t.Error(errors.New("incorrect number of records were updated"))
	}
}

func TestUpsert(t *testing.T) {
	input := UpsertInput{
		Record:   map[string]interface{}{"someField": 1, "field2": "string val", "field3": true, "addedField": "run2"},
		RecordId: "alexis2.test",
	}
	input.ServiceAccountJson = os.Getenv("FIREBASE_JSON")
	input.CollectionPath = "apps"

	step := Upsert{}
	_, err := step.execute(input, "runString")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
}

func getQueryParams() []QueryParam {
	query := QueryParam{
		FieldName:  "synced",
		Operator:   "==",
		FieldValue: true,
	}
	return []QueryParam{query}
}

func createEngine() step.Engine {
	environment := core.Environment{
		Project:  "unit_testing",
		Workflow: "fake_workflow",
		RunId:    "runString",
		Debug:    true,
	}
	return step.GetEngine(&environment)
}

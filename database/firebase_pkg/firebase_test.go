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

func TestQuery(t *testing.T) {
	input := QueryInput{
		QueryParameters: getQueryParams(),
	}
	input.StorageBucket = os.Getenv("FIREBASE_STORAGE_BUCKET")
	input.ServiceAccountJson = os.Getenv("FIREBASE_JSON")
	input.CollectionPath = "apps"

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
		Workflow: "fake_workflow_two",
	}
	input.StorageBucket = os.Getenv("FIREBASE_STORAGE_BUCKET")
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
		UpdateFields: map[string]interface{}{"newField": "newStringVal"},
		MaxUpdateSize: 1,
	}
	input.StorageBucket = os.Getenv("FIREBASE_STORAGE_BUCKET")
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
		Record: map[string]interface{}{"someField": 1, "field2": "string val", "field3": true, "addedField": "run2"},
		RecordId: "alexis2.test",
	}
	input.StorageBucket = os.Getenv("FIREBASE_STORAGE_BUCKET")
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
		FieldName:  "findMe",
		Operator:   "==",
		FieldValue: "hello",
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

package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"google.golang.org/api/option"
	"io/ioutil"
)

type FirebaseInput struct {
	StorageBucket      string
	ServiceAccountJson string
	CollectionPath     string
}

type QueryParam struct {
	FieldName  string
	Operator   string
	FieldValue string
}

type QueryOutput struct {
	Records []map[string]interface{}
}

func GetFirebaseAppFromConfig(serviceJson, storageBucket, runId string) (*firebase.App, error) {
	ctx := context.Background()
	tmpFile, err := ioutil.TempFile("", fmt.Sprintf("%sserviceAccount.json", runId))
	if err != nil {
		return nil, err
	}
	_, err = tmpFile.WriteString(serviceJson)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(tmpFile.Name())
	config := &firebase.Config{
		StorageBucket: storageBucket,
	}
	return firebase.NewApp(ctx, config, opt)
}

func QueryFirebase(app *firebase.App, coll string, queryParams []QueryParam) ([]*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	store, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	collection := store.Collection(coll)
	query := collection.Query
	for _, queryParam := range queryParams {
		query = query.Where(queryParam.FieldName, queryParam.Operator, queryParam.FieldValue)
	}
	docIterator := query.Documents(ctx)
	return docIterator.GetAll()
}

func ParseAndQueueRecords(docs []*firestore.DocumentSnapshot, workflow string, engine *step.Engine) ([]map[string]interface{}, error) {
	records := make([]map[string]interface{}, len(docs))
	for idx, snap := range docs {
		var record map[string]interface{}
		err := snap.DataTo(&record)
		if err != nil {
			return records, err
		}
		if workflow != "" {
			err = engine.AddToQueue(workflow, record)
			if err != nil {
				return records, err
			}
		}
		records[idx] = record
	}

	return records, nil
}



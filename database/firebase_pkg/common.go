package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"google.golang.org/api/option"
)

type FirebaseInput struct {
	ServiceAccountJson string
	CollectionPath     string
}

type QueryParam struct {
	FieldName  string
	Operator   string
	FieldValue interface{}
}

type QueryOutput struct {
	Records []map[string]interface{}
}

func GetFirebaseAppFromConfig(serviceJson, runId string) (*firebase.App, error) {
	opt := option.WithCredentialsJSON([]byte(serviceJson))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func QueryFirebase(store *firestore.Client, coll string, queryParams []QueryParam) ([]*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	collection := store.Collection(coll)
	if collection == nil {
		return nil, fmt.Errorf("Invalid collection path: %s", coll)
	}
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
		record["_id"] = snap.Ref.ID
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

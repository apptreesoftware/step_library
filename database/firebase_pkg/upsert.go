package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
)

type UpsertInput struct {
	FirebaseInput
	Record   map[string]interface{}
	RecordId string
}

type Upsert struct {
}

func (Upsert) Name() string {
	return "upsert"
}

func (Upsert) Version() string {
	return "1.0"
}

func (u Upsert) Execute(in step.Context) (interface{}, error) {
	var input UpsertInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return u.execute(input, in.Environment().RunId)
}

func (Upsert) execute(input UpsertInput, runId string) (interface{}, error) {
	app, err := GetFirebaseAppFromConfig(input.ServiceAccountJson, runId)
	if err != nil {
		return &UpdateDocumentOutput{}, err
	}

	ctx := context.Background()
	store, err := app.Firestore(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("Unable to connect to firestore: %w", err)
	}
	defer store.Close()

	collection := store.Collection(input.CollectionPath)
	doc := collection.Doc(input.RecordId)
	_, err = doc.Set(ctx, input.Record)
	return nil, err
}

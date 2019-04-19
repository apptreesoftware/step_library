package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type UpdateDocumentInput struct {
	FirebaseInput
	QueryParameters []QueryParam
	UpdateFields    map[string]interface{}
	MaxUpdateSize   int
}

type UpdateDocumentOutput struct {
	RecordsUpdated int
}

type UpdateDocument struct {
}

func (UpdateDocument) Name() string {
	return "update_doc"
}

func (UpdateDocument) Version() string {
	return "1.0"
}

func (u UpdateDocument) Execute(in step.Context) (interface{}, error) {
	var input UpdateDocumentInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return u.execute(input, in.Environment().RunId)
}

func (UpdateDocument) execute(input UpdateDocumentInput, runId string) (*UpdateDocumentOutput, error) {
	app, err := GetFirebaseAppFromConfig(input.ServiceAccountJson, input.StorageBucket, runId)
	if err != nil {
		return &UpdateDocumentOutput{}, err
	}

	snaps, err := QueryFirebase(app, input.CollectionPath, input.QueryParameters)
	if err != nil {
		return &UpdateDocumentOutput{}, err
	}
	if input.MaxUpdateSize != 0 && len(snaps) > input.MaxUpdateSize {
		return &UpdateDocumentOutput{}, errors.New("query returned more records than allowed per max update size in request")
	}

	updates := make([]firestore.Update, 0)
	for fieldPath, value := range input.UpdateFields {
		update := firestore.Update{
			Path:  fieldPath,
			Value: value,
		}
		updates = append(updates, update)
	}

	ctx := context.Background()
	for _, snap := range snaps {
		_, err := snap.Ref.Update(ctx, updates)
		if err != nil {
			return &UpdateDocumentOutput{}, err
		}
	}

	return &UpdateDocumentOutput{RecordsUpdated: len(snaps)}, nil
}

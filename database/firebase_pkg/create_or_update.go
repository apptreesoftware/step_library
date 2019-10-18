package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type CreateOrUpdate struct {
}

func (CreateOrUpdate) Name() string {
	return "create_or_update"
}

func (CreateOrUpdate) Version() string {
	return "1.0"
}

func (u CreateOrUpdate) Execute(in step.Context) (interface{}, error) {
	var input UpsertInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return u.execute(input, in.Environment().RunId)
}

func (CreateOrUpdate) execute(input UpsertInput, runId string) (interface{}, error) {
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
	_, err = doc.Get(ctx)
	if err != nil && grpc.Code(err) == codes.NotFound {
		_, err = doc.Set(ctx, input.Record)
	} else {
		updates := make([]firestore.Update, 0)
		for k, v := range input.Record {
			update := firestore.Update{
				Path:  k,
				Value: v,
			}
			updates = append(updates, update)
		}
		_, err = doc.Update(ctx, updates)
	}
	return nil, err
}

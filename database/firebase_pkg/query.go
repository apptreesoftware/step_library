package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
)

type QueryInput struct {
	FirebaseInput
	QueryParameters []QueryParam
}

type Query struct {
}

func (Query) Name() string {
	return "query"
}

func (Query) Version() string {
	return "1.0"
}

func (q Query) Execute(ctx step.Context) (interface{}, error) {
	input := QueryInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return q.execute(input, ctx.Environment().RunId)
}

func (Query) execute(input QueryInput, runId string) (*QueryOutput, error) {
	app, err := GetFirebaseAppFromConfig(input.ServiceAccountJson, runId)
	if err != nil {
		return &QueryOutput{}, err
	}
	store, err := app.Firestore(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("Unable to connect to firestore: %w", err)
	}
	defer store.Close()

	snaps, err := QueryFirebase(store, input.CollectionPath, input.QueryParameters)
	if err != nil {
		return &QueryOutput{}, err
	}

	records, err := ParseAndQueueRecords(snaps, "", nil)

	return &QueryOutput{Records: records}, nil
}

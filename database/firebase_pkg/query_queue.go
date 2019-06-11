package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
)

type QueryQueueInput struct {
	FirebaseInput
	QueryParameters []QueryParam
	Workflow        string
}

type QueryAndQueue struct {
}

func (QueryAndQueue) Name() string {
	return "query_and_queue"
}

func (QueryAndQueue) Version() string {
	return "1.0"
}

func (q QueryAndQueue) Execute(ctx step.Context) (interface{}, error) {
	var input QueryQueueInput
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	return q.execute(input, ctx.Environment().RunId, ctx.Engine())
}

func (QueryAndQueue) execute(input QueryQueueInput, runId string, engine step.Engine) (*QueryOutput, error) {
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

	records, err := ParseAndQueueRecords(snaps, input.Workflow, &engine)
	return &QueryOutput{Records: records}, err
}

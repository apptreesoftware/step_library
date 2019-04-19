package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
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
	app, err := GetFirebaseAppFromConfig(input.ServiceAccountJson, input.StorageBucket, runId)
	if err != nil {
		return &QueryOutput{}, err
	}

	snaps, err := QueryFirebase(app, input.CollectionPath, input.QueryParameters)
	if err != nil {
		return &QueryOutput{}, err
	}

	records, err := ParseAndQueueRecords(snaps, "", nil)

	return &QueryOutput{Records: records}, nil
}

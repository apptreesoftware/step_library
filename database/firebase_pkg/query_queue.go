package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

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
	app, err := GetFirebaseAppFromConfig(input.ServiceAccountJson, input.StorageBucket, runId)
	if err != nil {
		return &QueryOutput{}, err
	}

	snaps, err := QueryFirebase(app, input.CollectionPath, input.QueryParameters)
	if err != nil {
		return &QueryOutput{}, err
	}

	records, err := ParseAndQueueRecords(snaps, input.Workflow, &engine)
	return &QueryOutput{Records: records}, err
}

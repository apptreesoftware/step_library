package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type CachePush struct {
}

func (CachePush) Name() string {
	return "push"
}

func (CachePush) Version() string {
	return "1.0"
}

func (CachePush) Execute(in step.Context) (interface{}, error) {
	input := pushInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	engine := in.Engine()
	err = engine.PutRecord(input.Id, input.Record, map[string]interface{}{}, input.CacheName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type pushInput struct {
	Record    map[string]interface{}
	Id        string
	CacheName string
}

type CachePull struct {
}

func (CachePull) Name() string {
	return "pull"
}

func (CachePull) Version() string {
	return "1.0"
}

func (CachePull) Execute(in step.Context) (interface{}, error) {
	input := pullInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	engine := in.Engine()
	rec := map[string]interface{}{}
	found, _, err := engine.PullRecord(input.Id, &rec, input.CacheName)
	if err != nil {
		return nil, err
	}
	if !found {
		return pullOutput{Found: false}, nil
	}
	return pullOutput{
		Record: rec,
		Found:  true,
	}, nil
}

type pullInput struct {
	Id        string
	CacheName string
}

type pullOutput struct {
	Record map[string]interface{}
	Found  bool
}

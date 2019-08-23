package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/xerrors"
)

type BatchWrite struct {
}

func (BatchWrite) Name() string {
	return "batch_write"
}

func (BatchWrite) Version() string {
	return "1.0"
}

func (s BatchWrite) Execute(ctx step.Context) (interface{}, error) {
	var input BatchWriteInput
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return s.execute(input)
}

func (BatchWrite) execute(input BatchWriteInput) (*BatchWriteOutput, error) {
	srv, err := GetSheetsService(input.InputBase, false)
	if err != nil {
		return nil, err
	}

	return UpdateRecordsBatch(input.BatchBase, srv, input.Records)
}

type CacheBatchWrite struct {

}

func (CacheBatchWrite) Name() string {
	return "cache_batch_write"
}

func (CacheBatchWrite) Version() string {
	return "1.0"
}

func (c CacheBatchWrite) Execute(in step.Context) (interface{}, error) {
	var input CacheBatchWriteInput
	err := in.BindInputs(&input)
	if err != nil {
		return nil, xerrors.Errorf("Unable to read inputs: %v", err)
	}

	return c.execute(input, in.Engine())
}

func (CacheBatchWrite) execute(input CacheBatchWriteInput, engine step.Engine) (*BatchWriteOutput, error) {
	cacheRecords, err := engine.Find(input.Filter, input.CacheName, 0)
	if err != nil {
		return nil, xerrors.Errorf("Unable to fetch records from cache: %v", err)
	}

	srv, err := GetSheetsService(input.InputBase, false)
	if err != nil {
		return nil, err
	}

	records := make([]map[string]interface{}, len(cacheRecords))
	for idx, record := range cacheRecords {
		var recordMap map[string]interface{}
		err = bson.Unmarshal(record.Record, &recordMap)
		if err != nil {
			return nil, xerrors.Errorf("Unable to parse cache record: %v", err)
		}
		records[idx] = recordMap
	}

	return UpdateRecordsBatch(input.BatchBase, srv, records)
}

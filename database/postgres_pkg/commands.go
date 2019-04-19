package main

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library/database/db_common"
	_ "github.com/lib/pq"
)

type Query struct {
}

func (Query) Name() string {
	return "query"
}

func (Query) Version() string {
	return "1.0"
}

func (Query) Execute(ctx step.Context) (interface{}, error) {
	input := db_common.DatabaseCommand{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db_common.PerformQuery(db, input)
}

type QueryAndQueue struct {
}

func (QueryAndQueue) Name() string {
	return "query_and_queue"
}

func (QueryAndQueue) Version() string {
	return "1.0"
}

func (QueryAndQueue) Execute(ctx step.Context) (interface{}, error) {
	input := db_common.DatabaseCommandToQueue{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db_common.PerformQueryAndQueue(db, input, ctx.Engine())
}

type InsertBatch struct {
}

func (InsertBatch) Name() string {
	return "insert_batch"
}

func (InsertBatch) Version() string {
	return "1.0"
}

func (InsertBatch) Execute(ctx step.Context) (interface{}, error) {
	input := &db_common.InsertBatchCommand{}
	err := ctx.BindInputs(input)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	err = db_common.PerformInsertAll(db, input)
	return nil, err
}

type Execute struct {
}

func (Execute) Name() string {
	return "execute"
}

func (Execute) Version() string {
	return "1.0"
}

func (Execute) Execute(in step.Context) (interface{}, error) {
	input := db_common.DatabaseCommand{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	println(input.ConnectionString)
	db, err := sql.Open("postgres", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	_, err = db_common.ExecuteStatement(db, &input)
	return nil, err
}

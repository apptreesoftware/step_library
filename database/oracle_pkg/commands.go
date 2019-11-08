package main

import (
	"database/sql"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library/database/db_common"
	"golang.org/x/xerrors"
	_ "gopkg.in/goracle.v2"
	"time"
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
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db_common.PerformQuery(db, input)
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
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	err = db_common.PerformInsertAll(db, input)
	return nil, err
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
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db_common.PerformQueryAndQueue(db, input, ctx.Engine())
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
	debug := in.Environment().Debug
	if debug {
		println("Starting Step")
	}
	if err != nil {
		return nil, err
	}
	start := time.Now()
	if debug {
		println("Creating connection to", input.ConnectionString)
	}
	db, err := sql.Open("goracle", input.ConnectionString)
	dur := (time.Now().UnixNano() - start.UnixNano()) / int64(time.Millisecond)
	if debug {
		println(fmt.Sprintf("Connection established after %d ms", dur))
	}
	if err != nil {
		return nil, err
	}
	if debug {
		println("Executing statement", input.Sql)
	}
	start = time.Now()
	_, err = db_common.ExecuteStatement(db, &input)
	dur = (time.Now().UnixNano() - start.UnixNano()) / int64(time.Millisecond)
	if debug {
		println(fmt.Sprintf("ExecuteStatement took %d ms", dur))
	}
	return nil, err
}

type ImportCSV struct {
}

func (i ImportCSV) Name() string {
	return "import_csv"
}

func (i ImportCSV) Version() string {
	return "1.0"
}

func (i ImportCSV) Execute(in step.Context) (interface{}, error) {
	var input db_common.ImportCSVCommand
	err := in.BindInputs(&input)
	if err != nil {
		return nil, xerrors.Errorf("Unable to bind inputs: %w", err)
	}
	db, err := sql.Open("goracle", input.ConnectionString)
	if err != nil {
		return nil, xerrors.Errorf("Could not connect to database: %w", err)
	}
	return db_common.ImportCSV(db, "oracle", &input)
}

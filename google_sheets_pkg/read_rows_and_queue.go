package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"workflow/engine"
)

type ReadRowsAndQueue struct {}

func (ReadRowsAndQueue) Name() string {
	return "read_rows_and_queue"
}

func (ReadRowsAndQueue) Version() string {
	return "1.0"
}

func (s ReadRowsAndQueue) Execute(ctx step.Context)  (interface{}, error) {
	input := ReadRowsAndQueueInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	err = s.execute(input, ctx)
	return nil, err
}

func (ReadRowsAndQueue) execute(input ReadRowsAndQueueInput, ctx step.Context) error {
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), spreadsheet.Scope)
	if err != nil {
		return err
	}
	client := conf.Client(context.Background())
	service := spreadsheet.NewServiceWithClient(client)
	gsheet, err := service.FetchSpreadsheet(input.SpreadsheetId)
	if err != nil {
		return err
	}
	sheet, err := gsheet.SheetByIndex(input.SheetIndex)
	if err != nil {
		return err
	}
	wfEngine := ctx.Engine()

	for _, row := range sheet.Rows {
		rowMap := map[string]interface{}{}
		for colIndex, name := range input.Fields {
			rowMap[name] = row[colIndex].Value
		}
		err = wfEngine.AddToQueue(input.Workflow, rowMap)
		if err != nil {
			return err
		}
	}
	return nil
}

type ReadRowsAndQueueInput struct {
	SpreadsheetId string
	SheetIndex uint
	Credentials string
	Fields []string
	Workflow string
}
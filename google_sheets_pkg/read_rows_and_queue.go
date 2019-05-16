package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type ReadRowsAndQueue struct{}

func (ReadRowsAndQueue) Name() string {
	return "read_rows_and_queue"
}

func (ReadRowsAndQueue) Version() string {
	return "1.0"
}

func (s ReadRowsAndQueue) Execute(ctx step.Context) (interface{}, error) {
	input := ReadRowsAndQueueInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	records, err := s.readSheet(input)
	err = s.queueWorkflows(input, records, ctx)
	return nil, err
}

func (ReadRowsAndQueue) readSheet(input ReadRowsAndQueueInput) ([]map[string]interface{}, error) {
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), spreadsheet.Scope)
	if err != nil {
		return nil, err
	}
	client := conf.Client(context.Background())
	service := spreadsheet.NewServiceWithClient(client)
	gsheet, err := service.FetchSpreadsheet(input.SpreadsheetId)
	if err != nil {
		return nil, err
	}
	sheet, err := gsheet.SheetByIndex(input.SheetIndex)
	if err != nil {
		return nil, err
	}

	var rows []map[string]interface{}
	for _, row := range sheet.Rows {
		rowMap := map[string]interface{}{}
		for colIndex, name := range input.Fields {
			rowMap[name] = row[colIndex].Value
		}
		rows = append(rows, rowMap)
	}

	return rows, nil
}

func (ReadRowsAndQueue) queueWorkflows(input ReadRowsAndQueueInput, records []map[string]interface{}, ctx step.Context) error {
	wfEngine := ctx.Engine()
	for _, record := range records {
		err := wfEngine.AddToQueue(input.Workflow, record)
		if err != nil {
			return err
		}
	}
	return nil
}

type ReadRowsAndQueueInput struct {
	SpreadsheetId string
	SheetIndex    uint
	Credentials   string
	Fields        []string
	Workflow      string
}

package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"google.golang.org/api/sheets/v4"
	"log"
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
	srv, err := ValidateInputAndGetConf(input.InputBase, false)
	if err != nil {
		return nil, err
	}

	created := 0
	updated := 0
	appendCounter := 1

	existingRecordMap := make(map[string]string)

	if input.ClearSheet {
		err := ClearSheet(input.InputBase, srv)
		if err != nil {
			return nil, err
		}
	} else {
		sheet, err := GetSheet(input.InputBase, srv)
		if err != nil {
			return nil, err
		}

		data := sheet.Data
		defaultRange := data[0]
		rows := defaultRange.RowData
		appendCounter = len(rows) + 1

		if input.Update {
			existingRecordMap = getExistingRecordMap(sheet, input.MatchColumn)
		}

	}

	for _, record := range input.Records {
		idValue := fmt.Sprintf("%v", record[input.MatchColumn])
		rowIndex := existingRecordMap[idValue]
		if !input.Update || rowIndex == "" {
			rowIndex = fmt.Sprintf("A%d", appendCounter)
			appendCounter++
			created++
		} else {
			updated++
		}
		var vr sheets.ValueRange
		vr.Values = append(vr.Values, record)
		_, err := srv.Spreadsheets.Values.Update(input.SpreadsheetId, rowIndex, &vr).ValueInputOption("RAW").Do()
		if err != nil {
			log.Fatalf("Unable to update data from sheet. %v", err)
		}
	}
	return &BatchWriteOutput{RecordsUpdated: updated, RecordsCreated: created}, nil
}

func getExistingRecordMap(sheet *sheets.Sheet, matchColumn int) map[string]string {
	matchValMap := make(map[string]string)
	data := sheet.Data
	defaultRange := data[0]
	rows := defaultRange.RowData
	for rowIdx, row := range rows {
		values := row.Values
		if matchColumn > len(values) {
			continue
		}
		matchVal := values[matchColumn]
		if matchVal != nil {
			value := matchVal.FormattedValue
			matchValMap[value] = fmt.Sprintf("A%d", rowIdx+1)
		}
	}
	return matchValMap
}

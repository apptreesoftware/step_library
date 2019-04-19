package main

import (
	"context"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type WriteToSheet struct{}

func (WriteToSheet) Name() string {
	return "write"
}

func (WriteToSheet) Version() string {
	return "1.0"
}

func (w WriteToSheet) Execute(ctx step.Context) (interface{}, error) {
	input := WriteToSheetInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	err = w.execute(input)
	return nil, err
}

func (WriteToSheet) execute(input WriteToSheetInput) error {
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), spreadsheet.Scope)
	if err != nil {
		return err
	}
	client := conf.Client(context.Background())
	service := spreadsheet.NewServiceWithClient(client)
	googleSheet, err := service.FetchSpreadsheet(input.SpreadsheetId)
	if err != nil {
		return err
	}
	sheet, err := googleSheet.SheetByIndex(input.SheetIndex)
	if err != nil {
		return err
	}

	isUpdate := false
	// If user put `MatchValue` input AND Document isn't brand new
	// sheet.Rows refers to the number of rows _with_ data in the Sheet
	// blank rows aren't counted.
	if input.MatchValue != "" && len(sheet.Rows) > 0 {
		isUpdate = true
	}

	// find the row to be updated, if it exists
	recordRowIdxToUpdate := -1
	if isUpdate {
		for idx, row := range sheet.Rows {
			if row[input.MatchColumn].Value == input.MatchValue {
				recordRowIdxToUpdate = idx
				break
			}
		}
		if recordRowIdxToUpdate == -1 {
			fmt.Printf("Record with value %s in column %v doesn't exist. Adding the record to the end of the sheet", input.MatchValue, input.MatchColumn)
			isUpdate = false
		}
	}

	// Finally, if this is an update form, update Row @ recordRowIdxToUpdate
	data := sheet.Properties.GridProperties
	var newRow int
	if isUpdate {
		newRow = recordRowIdxToUpdate
		// If it isn't an update form,
		// OR the record wasn't found
		// add the record to the end
	} else {
		newRow = int(data.RowCount)
	}

	// Update the record
	for colIdx, cellVal := range input.Cells {
		sheet.Update(int(newRow), colIdx, fmt.Sprintf("%v", cellVal))
	}
	err = sheet.Synchronize()
	return err
}

type WriteToSheetInput struct {
	SpreadsheetId string
	SheetIndex    uint
	Credentials   string
	Cells         []interface{}
	MatchColumn   int
	MatchValue    string
}

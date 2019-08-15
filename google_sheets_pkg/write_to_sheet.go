package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
	"google.golang.org/api/sheets/v4"
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
	srv, err := GetSheetsService(input.InputBase, false)
	if err != nil {
		return err
	}

	sheet, err := GetSheet(input.InputBase, srv)
	if err != nil {
		return err
	}

	data := sheet.Data
	defaultRange := data[0]
	rows := defaultRange.RowData

	updating := false
	appendCounter := len(rows) +1
	for idx, row := range rows {
		values := row.Values
		if input.MatchColumn > len(values) {
			continue
		}
		matchVal := values[input.MatchColumn]
		if matchVal != nil {
			value := matchVal.FormattedValue
			if value == input.MatchValue {
				updating = true
				appendCounter = idx+1
				break
			}
		}
	}

	var vr sheets.ValueRange
	vr.Values = append(vr.Values, input.Cells)
	rowIndex := fmt.Sprintf("A%d", appendCounter)

	if updating {
		_, err = srv.Spreadsheets.Values.Update(input.SpreadsheetId, rowIndex, &vr).ValueInputOption("RAW").Do()
	} else {
		_, err = srv.Spreadsheets.Values.Append(input.SpreadsheetId, rowIndex, &vr).ValueInputOption("RAW").Do()
	}
	if err != nil {
		return xerrors.Errorf("Error encountered saving row: %v", err)
	}
	return nil
}



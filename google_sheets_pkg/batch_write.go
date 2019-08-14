package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
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
	srv, err := GetSheetsService(input.InputBase, false)
	if err != nil {
		return nil, err
	}

	created := 0
	updated := 0
	appendCounter := 1
	rows := make([]*sheets.RowData, 0)
	maxColumns := input.GetHighestFieldIndex()

	sheet, err := GetSheet(input.InputBase, srv)
	if err != nil {
		return nil, err
	}

	existingRecordMap := make(map[string]*DataHelper)

	if input.ClearSheet {
		err := ClearSheet(input.InputBase, srv)
		if err != nil {
			return nil, err
		}
	} else {
		matchIndex, ok := input.GetIndexForId()
		if !ok {
			return nil, xerrors.Errorf("Unable to find column index for ID field. It is missing from fields list")
		}

		data := sheet.Data
		defaultRange := data[0]
		rows = defaultRange.RowData
		appendCounter = len(rows) + 1

		if input.Update {
			existingRecordMap = getExistingRecordMap(sheet, matchIndex)
		}

	}

	for _, record := range input.Records {
		updating := false
		idValue := fmt.Sprintf("%v", record[input.MatchColumn])
		existingData := existingRecordMap[idValue]
		var rowIndex string
		if existingData == nil {
			rowIndex = fmt.Sprintf("A%d", appendCounter)
			created++
			appendCounter++
		} else {
			rowIndex = existingData.StartCell
			updated++
			updating = true
		}

		vr := createRowFromRecord(record, input.Fields, existingData, maxColumns)
		if updating {
			_, err = srv.Spreadsheets.Values.Update(input.SpreadsheetId, rowIndex, &vr).ValueInputOption("RAW").Do()
		} else {
			_, err = srv.Spreadsheets.Values.Append(input.SpreadsheetId, rowIndex, &vr).ValueInputOption("RAW").Do()
		}
		if err != nil {
			log.Fatalf("Unable to update data from sheet. %v", err)
		}
	}
	return &BatchWriteOutput{RecordsUpdated: updated, RecordsCreated: created}, nil
}

func getExistingRecordMap(sheet *sheets.Sheet, matchColumn int) map[string]*DataHelper {
	matchValMap := make(map[string]*DataHelper)
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
			matchValMap[value] = &DataHelper{
				Data:      row,
				StartCell: fmt.Sprintf("A%d", rowIdx+1),
			}
		}
	}
	return matchValMap
}

func createRowFromRecord(record map[string]interface{}, fields map[int]string, existing *DataHelper,
	maxFields int) sheets.ValueRange {
	var vr sheets.ValueRange

	val := make([]interface{}, maxFields + 1)
	for i := 0; i <= maxFields; i++ {
		fieldName := fields[i]
		var value interface{}
		if (fieldName == "" || record[fieldName] == nil) && existing != nil {
			value = existing.Data.Values[i].FormattedValue
		} else {
			value = record[fieldName]
		}
		if value != nil {
			val[i] = fmt.Sprintf("%v", value)
		}
	}
	vr.Values = append(vr.Values, val)

	return vr
}

package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"strings"
)

func GetSheetsService(input InputBase, readOnly bool) (*sheets.Service, error) {
	if input.Credentials == "" {
		return nil, errors.New("No credentials were provided. Please provide your google service account key using the `Credentials` input")
	}

	url := "https://www.googleapis.com/auth/spreadsheets"
	if readOnly {
		url = url + ".readonly"
	}
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), url)
	if err != nil {
		return nil, xerrors.Errorf("Unable to create spreadsheet config: %w", err)
	}

	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(conf.Client(ctx)))
	if err != nil {
		return nil, xerrors.Errorf("Unable to create spreadsheet service: %v", err)
	}
	return srv, nil
}

func GetSheet(input InputBase, srv *sheets.Service) (*sheets.Sheet, error) {
	spreadsheet, err := srv.Spreadsheets.
		Get(input.SpreadsheetId).
		IncludeGridData(true).
		Do()
	if err != nil {
		return nil, xerrors.Errorf("Unable to fetch spreadsheet %s: %v",
			input.SpreadsheetId, err)
	}

	if len(input.SheetName) > 0 {
		var names []string
		for _, sheet := range spreadsheet.Sheets {
			if sheet.Properties.Title == input.SheetName {
				return sheet, nil
			}
			names = append(names, sheet.Properties.Title)
		}

		return nil, xerrors.Errorf("SheetName %s not found. Sheets found: %s", input.SheetName, strings.Join(names, ","))
	}
	if len(spreadsheet.Sheets) <= int(input.SheetIndex) {
		return nil, xerrors.Errorf("SheetIndex is out of range. The maximum sheet index for this spreadsheet is %d", len(spreadsheet.Sheets))
	}
	return spreadsheet.Sheets[input.SheetIndex], nil
}

func ClearSheet(input InputBase, srv *sheets.Service) error {
	clearReq := sheets.BatchClearValuesRequest{
		Ranges: []string{"A2:ZZ"},
	}
	_, err := srv.Spreadsheets.Values.BatchClear(input.SpreadsheetId, &clearReq).Do()
	if err != nil {
		return xerrors.Errorf("Unable to clear sheet: %w", err)
	}
	return nil
}

func GetRowsFromSheet(sheet *sheets.Sheet, matchColumn int) map[string]*DataHelper {
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

func UpdateRecordsBatch(input BatchBase, srv *sheets.Service, records []map[string]interface{}) (*BatchWriteOutput, error) {
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
			existingRecordMap = GetRowsFromSheet(sheet, matchIndex)
		}

	}

	for _, record := range records {
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

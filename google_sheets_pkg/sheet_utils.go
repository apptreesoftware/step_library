package main

import (
	"context"
	"errors"
	"golang.org/x/oauth2/google"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"strings"
)

func ValidateInputAndGetConf(input InputBase, readOnly bool) (*sheets.Service, error) {
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
	return srv, err
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

func ParseFields(sheet *sheets.Sheet, fields []string, startRow int) []map[string]string {
	var rowValues []map[string]string
	data := sheet.Data
	defaultRange := data[0]
	rows := defaultRange.RowData
	rows = rows[startRow:]
	for _, row := range rows {
		rowMap := map[string]string{}
		values := row.Values
		for columnIndex, cellValue := range values {
			if columnIndex >= len(fields) {
				break
			}
			value := cellValue.FormattedValue
			fieldName := fields[columnIndex]
			rowMap[fieldName] = value
		}
		rowValues = append(rowValues, rowMap)
	}
	return rowValues
}

func ParseRow(sheet *sheets.Sheet, rowIndex int) []string {
	var rowData []string
	data := sheet.Data
	defaultRange := data[0]
	rows := defaultRange.RowData
	lastDataIndex := 0
	row := rows[rowIndex]
	for columnIndex, cellValue := range row.Values {
		value := cellValue.FormattedValue
		rowData = append(rowData, value)
		if len(value) != 0 {
			lastDataIndex = columnIndex
		}
	}
	return rowData[0:lastDataIndex]
}

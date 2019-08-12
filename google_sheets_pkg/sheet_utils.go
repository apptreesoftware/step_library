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

func ValidateInputAndGetConf(input InputBase) (*sheets.Service, error) {
	if input.Credentials == "" {
		return nil, errors.New("No credentials were provided. Please provide your google service account key using the `Credentials` input")
	}

	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), "https://www.googleapis.com/auth/spreadsheets.readonly")
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

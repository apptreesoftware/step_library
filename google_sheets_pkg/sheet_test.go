package main

import (
	"os"
	"testing"
)

func TestReadSheetUseRowAsFields(t *testing.T) {
	sheetId := os.Getenv("SHEET_ID")
	authToken := os.Getenv("SHEET_AUTH_TOKEN")

	input := ReadSheetInput{
		SpreadsheetId:       sheetId,
		SheetIndex:          0,
		Credentials:         authToken,
		UseFirstRowAsFields: true,
	}

	step := ReadSheet{}
	output, err := step.execute(input)
	if err != nil {
		t.Error(err)
		return
	}
	if len(output.Rows) == 0 {
		t.Log("0 rows returned")
		t.Fail()
		return
	}

	if output.Rows[0]["Config Parameter"] != "NotificationTitle" {
		t.Log("First now did not contain NotificationTitle")
		t.Fail()
	}
}

func TestGetSheetByName(t *testing.T) {
	sheetId := os.Getenv("SHEET_ID")
	authToken := os.Getenv("SHEET_AUTH_TOKEN")

	input := ReadSheetInput{
		SpreadsheetId: sheetId,
		SheetName:     "Contacts",
		Credentials:   authToken,
		Fields:        []string{"Value"},
	}

	step := ReadSheet{}
	output, err := step.execute(input)
	if err != nil {
		t.Error(err)
		return
	}
	if len(output.Rows) == 0 {
		t.Log("0 rows returned")
		t.Fail()
		return
	}

	if output.Rows[0]["Value"] != "ContactID" {
		t.Log("First now did not contain NotificationTitle")
		t.Fail()
	}
}

func TestReadSheetFieldsInput(t *testing.T) {
	sheetId := os.Getenv("SHEET_ID")
	authToken := os.Getenv("SHEET_AUTH_TOKEN")

	input := ReadSheetInput{
		SpreadsheetId: sheetId,
		SheetIndex:    0,
		Credentials:   authToken,
		Fields: []string{
			"ConfigParam",
			"ConfigParamValue",
		},
	}

	step := ReadSheet{}
	output, err := step.execute(input)
	if err != nil {
		t.Error(err)
		return
	}
	if len(output.Rows) == 0 {
		t.Log("0 rows returned")
		t.Fail()
		return
	}

	if output.Rows[1]["ConfigParam"] != "NotificationTitle" {
		t.Log("First now did not contain NotificationTitle")
		t.Fail()
	}
}

func TestUseFirstFieldsMutuallyExclusive(t *testing.T) {
	input := ReadSheetInput{
		Fields: []string{
			"ConfigParam",
			"ConfigParamValue",
		},
		UseFirstRowAsFields: true,
	}
	step := ReadSheet{}
	_, err := step.execute(input)
	if !(err != nil && err.Error() == "If you specify UseFirstRowAsFields, then you can not specify Fields.") {
		t.Fail()
	}
}

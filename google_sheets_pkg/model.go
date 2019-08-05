package main

type InputBase struct {
	SpreadsheetId string
	SheetIndex    uint
	SheetName     string
	Credentials   string
}

type ReadSheetInput struct {
	InputBase
	Fields              []string
	UseFirstRowAsFields bool
}

type ReadSheetOutput struct {
	SheetName string
	Rows      []map[string]string
}

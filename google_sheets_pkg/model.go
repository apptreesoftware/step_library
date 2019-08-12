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

type BatchWriteInput struct {
	InputBase
	Overwrite    bool
	Records      []interface{}
	MatchColumns map[string]string
	IdColumn     string
}

type BatchWriteOutput struct {
	NumberProcessed int
}

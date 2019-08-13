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
	Records     [][]interface{}
	MatchColumn int
	Update      bool
	ClearSheet  bool
}

type BatchWriteOutput struct {
	RecordsUpdated int
	RecordsCreated int
}

package main

import "google.golang.org/api/sheets/v4"

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
	Records     []map[string]interface{}
	Fields      map[int]string
	MatchColumn string
	Update      bool
	ClearSheet  bool
}

func (b BatchWriteInput) GetHighestFieldIndex() int {
	highestIndex := 0
	for key, _ := range b.Fields {
		if key > highestIndex {
			highestIndex = key
		}
	}
	return highestIndex
}

func (b BatchWriteInput) GetIndexForId() (index int, ok bool) {
	for k, val := range b.Fields {
		if val == b.MatchColumn {
			index = k
			ok = true
			return
		}
	}
	return
}

type BatchWriteOutput struct {
	RecordsUpdated int
	RecordsCreated int
}

type DataHelper struct {
	Data      *sheets.RowData
	StartCell string
}

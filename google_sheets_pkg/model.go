package main

import "google.golang.org/api/sheets/v4"

type InputBase struct {
	SpreadsheetId string
	SheetIndex    uint
	SheetName     string
	Credentials   string
}

type BatchBase struct {
	InputBase
	Fields      map[int]string
	MatchColumn string
	Update      bool
	ClearSheet  bool
}

type ReadSheetInput struct {
	InputBase
	Fields              []string
	UseFirstRowAsFields bool
	Ranges              []string
}

type ReadSheetOutput struct {
	SheetName string
	Rows      []map[string]string
}

type BatchWriteInput struct {
	BatchBase
	Records []map[string]interface{}
}

func (b BatchBase) GetHighestFieldIndex() int {
	highestIndex := 0
	for key, _ := range b.Fields {
		if key > highestIndex {
			highestIndex = key
		}
	}
	return highestIndex
}

func (b BatchBase) GetIndexForId() (index int, ok bool) {
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

type CacheBatchWriteInput struct {
	BatchBase
	CacheName string
	Filter    map[string]interface{}
}

type WriteToSheetInput struct {
	InputBase
	Cells         []interface{}
	MatchColumn   int
	MatchValue    string
}

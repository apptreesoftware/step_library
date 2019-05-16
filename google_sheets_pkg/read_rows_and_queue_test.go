package main

import (
	"fmt"
	"testing"
)

func TestReadRowsAndQueue_readSheet(t *testing.T) {
	cred := `` // enter your google credentials to test

	input := ReadRowsAndQueueInput{
		SpreadsheetId: "1Oq8k-7oxQ8ighdT_UrBLPct5wX2gTnKxyLfTQOdzxjc",
		SheetIndex:    0,
		Credentials:   cred,
		Fields:        []string{"FirstName", "LastName", "BirthYear", "Coolness"},
		Workflow:      "my_workflow",
	}

	readRows := ReadRowsAndQueue{}
	output, err := readRows.readSheet(input)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for i, record := range output {
		fmt.Printf("%v : %v", i, record)
	}
}

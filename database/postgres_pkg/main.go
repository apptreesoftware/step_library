package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library/database/db_common"
)

func main() {
	step.Register(Query{})
	step.Register(QueryAndQueue{})
	step.Register(InsertBatch{})
	step.Register(Execute{})
	step.Register(db_common.MockExecute{})
	step.Run()
}

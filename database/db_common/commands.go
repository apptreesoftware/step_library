package db_common

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func PerformQuery(db *sql.DB, command DatabaseCommand) (interface{}, error) {
	rows, err := db.Query(command.Sql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	results := []map[string]interface{}{}
	for rows.Next() {
		rowMap, err := ScanIntoMap(rows, cols)
		if err != nil {
			return nil, err
		}
		results = append(results, rowMap)
	}
	output := &RowOutput{
		Results: results,
	}
	return output, nil
}

func PerformQueryAndQueue(db *sql.DB, command DatabaseCommandToQueue, engine step.Engine) (interface{}, error) {
	rows, err := db.Query(command.Sql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		rowMap, err := ScanIntoMap(rows, cols)
		if err != nil {
			return nil, err
		}
		err = engine.AddToQueue(command.Workflow, rowMap)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func PerformInsertAll(db *sql.DB, command *InsertBatchCommand) error {
	if len(command.Records) == 0 {
		return nil
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, rec := range command.Records {
		var rowValues []interface{}
		for _, fieldName := range command.Fields {
			value := rec[fieldName]
			rowValues = append(rowValues, value)
		}
		_, err := tx.Exec(command.Sql, rowValues...)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func ExecuteStatement(db *sql.DB, command *DatabaseCommand) (sql.Result, error) {
	return db.Exec(command.Sql)
}

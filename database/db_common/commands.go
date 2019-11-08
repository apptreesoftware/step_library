package db_common

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/xerrors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
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

func PerformQueryWithArgs(db *sql.DB, command DatabaseCommand, args ...interface{}) (interface{}, error) {
	rows, err := db.Query(command.Sql, args...)
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

func ImportCSV(db *sql.DB, dbType string, command *ImportCSVCommand) (*ImportCSVOutput, error) {
	csvFile, err := os.Open(command.FilePath)
	if err != nil {
		return nil, xerrors.Errorf("Unable to read file: %w", err)
	}
	start := time.Now()
	reader := csv.NewReader(csvFile)
	i := 0
	tx, err := db.Begin()
	if err != nil {
		return nil, xerrors.Errorf("Unable to begin transaction: %w", err)
	}
	defer tx.Rollback()

	fieldDescriptors, err := getFieldDescriptions(command.Fields)
	if err != nil {
		return nil, err
	}

	if command.ClearTable {
		_, err := tx.Exec(fmt.Sprintf("DELETE FROM %s", command.TableName))
		if err != nil {
			return nil, xerrors.Errorf("Unable to execute delete command: %w", err)
		}
	}
	columns, valuePlaceholders := buildColumns(fieldDescriptors, dbType)
	sqlStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		command.TableName,
		columns,
		valuePlaceholders)
	println("Preparing SQL:", sqlStatement)
	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, xerrors.Errorf("Unable to prepare statement: %w", err)
	}
	for {
		line, err := reader.Read()
		if i == 0 && command.SkipFirst {
			i++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, xerrors.Errorf("Unable to read line %d: %w", i, err)
		}
		values, err := getValues(fieldDescriptors, line, i)
		if err != nil {
			return nil, err
		}
		_, err = stmt.Exec(values...)
		if err != nil {
			return nil, xerrors.Errorf("Unable to execute statement: %w", err)
		}
		println("Processing line", i)
		i++
	}
	err = tx.Commit()
	dur := time.Now().Unix() - start.Unix()
	if err != nil {
		return nil, xerrors.Errorf("Unable to commit transaction: %w", err)
	}
	return &ImportCSVOutput{
		RecordCount: i - 1,
		Duration:    fmt.Sprintf("%d seconds", dur)}, nil
}

func getValues(fieldDescriptors []CSVFieldDescription, line []string, lineNum int) ([]interface{}, error) {
	var vals []interface{}
	for _, field := range fieldDescriptors {
		if field.index >= len(line) {
			return nil, xerrors.Errorf("Invalid field specified - %s: %i. %i is beyond the number of columns in this csv")
		}
		data := line[field.index]
		switch field.columnType {
		case "text":
			vals = append(vals, data)
		case "integer":
			intVal, err := strconv.Atoi(data)
			if err != nil {
				return nil, xerrors.Errorf("Line: %d, Field: %s. Unable to convert value %s to an integer", lineNum, field.columnName, data)
			}
			vals = append(vals, intVal)
		}
	}
	return vals, nil
}

func buildColumns(fieldDescriptors []CSVFieldDescription, dbType string) (string, string) {
	var keys []string
	var values []string
	for i, field := range fieldDescriptors {
		keys = append(keys, field.columnName)
		switch dbType {
		case "oracle":
			values = append(values, fmt.Sprintf(":v%d", i+1))
		case "postgres":
			values = append(values, fmt.Sprintf("$%d", i+1))
		}
	}
	return strings.Join(keys, ","), strings.Join(values, ",")
}

func getFieldDescriptions(fields []string) ([]CSVFieldDescription, error) {
	var fieldDesc []CSVFieldDescription
	for _, field := range fields {
		components := strings.Split(field, ":")
		switch len(components) {
		case 2:
			columnName := components[0]
			fieldIndex, err := strconv.Atoi(components[1])
			if err != nil {
				return nil, xerrors.Errorf("Unable to parse column index for field %s. The valid format is `columnName:2`", columnName)
			}
			fieldDesc = append(fieldDesc, CSVFieldDescription{
				columnName: columnName,
				index:      fieldIndex,
				columnType: "text"})
		case 3:
			columnName := components[0]
			fieldIndex, err := strconv.Atoi(components[1])
			columnType := components[2]
			if err != nil {
				return nil, xerrors.Errorf("Unable to parse column index for field %s. The valid format is `columnName:2`", columnName)
			}
			fieldDesc = append(fieldDesc, CSVFieldDescription{
				columnName: columnName,
				index:      fieldIndex,
				columnType: columnType})
		}
	}
	return fieldDesc, nil
}

type CSVFieldDescription struct {
	columnName string
	index      int
	columnType string
}

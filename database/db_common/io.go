package db_common

import "database/sql"

type DatabaseCommand struct {
	ConnectionString string
	Sql              string
}

type DatabaseCommandToQueue struct {
	ConnectionString string
	Sql              string
	Workflow         string
}

type RowOutput struct {
	Results []map[string]interface{}
}

type InsertBatchCommand struct {
	DatabaseCommand
	Records []map[string]interface{}
	Fields  []string
}

type InsertCommand struct {
	DatabaseCommand
	Record map[string]interface{}
}

func ScanIntoMap(rows *sql.Rows, cols []string) (map[string]interface{}, error) {
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i, _ := range columns {
		columnPointers[i] = &columns[i]
	}

	// Scan the result into the column pointers...
	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	m := make(map[string]interface{})
	for i, colName := range cols {
		val := columnPointers[i].(*interface{})
		m[colName] = *val
	}
	return m, nil
}

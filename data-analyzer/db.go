package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

func extractData(db *sql.DB, table string, columns string) ([]map[string]interface{}, error) {

	query := fmt.Sprintf("SELECT %s FROM %s", columns, table)
	stmt, err := db.Prepare(query)
	if err!= nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err!= nil {
		return nil, err
	}
	defer rows.Close()


	var data []map[string]interface{}
	for rows.Next() {
		columns, err := rows.Columns()
		if err!= nil {
			return nil, err
		}
		row := make(map[string]interface{}, len(columns))
		for i, col := range columns {
			var val interface{}
			err = rows.Scan(&val)
			if err!= nil {
				return nil, err
			}
			row[col] = val
		}
		data = append(data, row)
	}
	return data, nil
}

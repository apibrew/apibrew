package hightest

import (
	"database/sql"
	"fmt"
)

func compareDatabases(db1 *sql.DB, db2 *sql.DB) error {
	// Get the list of tables in the first database
	rows1, err := db1.Query(`
        SELECT table_name 
        FROM information_schema.tables 
        WHERE table_schema = 'public'`)
	if err != nil {
		return err
	}
	defer rows1.Close()

	// Get the list of tables in the second database
	rows2, err := db2.Query(`
        SELECT table_name 
        FROM information_schema.tables 
        WHERE table_schema = 'public'`)
	if err != nil {
		return err
	}
	defer rows2.Close()

	// Compare the tables in the two databases
	for rows1.Next() {
		var tableName string
		rows1.Scan(&tableName)

		// Check if the second database has the same table
		var found bool
		for rows2.Next() {
			var tableName2 string
			rows2.Scan(&tableName2)
			if tableName == tableName2 {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Table %s not found in second database", tableName)
		}

		// Compare the data in the two tables
		rows1Data, err := db1.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			return err
		}
		defer rows1Data.Close()

		rows2Data, err := db2.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
		if err != nil {
			return err
		}
		defer rows2Data.Close()

		for rows1Data.Next() {
			// Get the row from the first database
			row1Data := make([]interface{}, 0)
			row1Pointers := make([]interface{}, 0)
			cols1, err := rows1Data.Columns()
			if err != nil {
				return err
			}
			for range cols1 {
				var data interface{}
				row1Data = append(row1Data, &data)
				row1Pointers = append(row1Pointers, &data)
			}
			rows1Data.Scan(row1Pointers...)

			// Check if the second database has the same row
			var found bool
			for rows2Data.Next() {
				row2Data := make([]interface{}, 0)
				row2Pointers := make([]interface{}, 0)
				cols2, err := rows2Data.Columns()
				if err != nil {
					return err
				}
				for range cols2 {
					var data interface{}
					row2Data = append(row2Data, &data)
					row2Pointers = append(row2Pointers, &data)
				}
				rows2Data.Scan(row2Pointers...)

				if fmt.Sprintf("%v", row1Data) == fmt.Sprintf("%v", row2Data) {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Data not found in table %s in second database", tableName)
			}
		}
	}

	return nil
}

package helpers

import (
	"database/sql"
	"fmt"
	"log"
)


func TableExists(db *sql.DB, tableName string) bool {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)
	var exists bool 
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking if table exists: %v", err)
	}
	return exists
}	
package main

import (
	"database/sql"
)

func OpenDb(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic("Couldn't create a DB")
	}
	return db
}

func ApplyBaseMigrations(db *sql.DB) {
	// To make sure all tables are present
	createFindingsTable(db)
	createToolCacheTable(db)
	// Key table to fill
	createCodebasesTable(db)
}

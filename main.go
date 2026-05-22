package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	InitConfig()

	_ = os.Remove(GetConfig().DbPath)
	db := OpenDb("sqlite", GetConfig().DbPath)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	ApplyBaseMigrations(db)

	f, err := os.Open(GetConfig().RepositoriesCsvAbsolutePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(f)

	repositories := ParseCsvToRepositories(f)
	codebases := CreateCodebases(GetConfig(), repositories, time.Now())
	InsertCodebases(db, codebases)
}

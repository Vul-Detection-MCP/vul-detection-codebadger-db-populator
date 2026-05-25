package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

func CreateRangeInclusive(start int, end int) []int {
	values := make([]int, end-start+1)
	for i := range values {
		values[i] = start + i
	}
	return values
}

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

	minPort, _ := strconv.Atoi(config.JoernMinPort)
	maxPort, _ := strconv.Atoi(config.JoernMaxPort)
	portsRange := CreateRangeInclusive(minPort, maxPort)

	repositories := ParseCsvToRepositories(f)
	codebases := CreateCodebases(GetConfig(), repositories, time.Now())
	InsertCodebases(db, codebases, portsRange)
}

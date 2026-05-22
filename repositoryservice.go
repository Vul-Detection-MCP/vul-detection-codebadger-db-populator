package main

import (
	"encoding/csv"
	"io"
	"log"
)

func ParseCsvToRepositories(reader io.Reader) []Repository {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = 4
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV.")
	}
	total := make([]Repository, len(records))
	for i, row := range records {
		repository := Repository{
			hash:     row[2],
			language: row[3],
		}
		total[i] = repository
	}

	return total
}

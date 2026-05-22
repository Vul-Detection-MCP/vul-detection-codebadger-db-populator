package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"path"
	"time"
)

const (
	DefaultStatus      = "ready"
	CpgBinName         = "cpg.bin"
	SelectAllStatement = "SELECT " +
		"hash, " +
		"source_path," +
		"language, " +
		"cpg_path, " +
		"metadata, " +
		"created_at " +
		"FROM codebases"
)

func InsertCodebases(db *sql.DB, codebases []Codebase) {
	stmt, err := db.Prepare("INSERT INTO codebases(" +
		"hash, " +
		"source_path, " +
		"language, " +
		"cpg_path, " +
		"metadata, " +
		"created_at, " +
		"last_accessed" +
		") VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("Can't prepare a statement.")
	}
	defer func(st *sql.Stmt) {
		_ = st.Close()
	}(stmt)

	for _, codebase := range codebases {
		jsonMetadataBytes, err := json.Marshal(codebase.metadata)
		if err != nil {
			log.Fatalf("Couldn't marshal: %#v\n", codebase.metadata)
		}
		jsonMetadata := string(jsonMetadataBytes)
		timestampFormatted := ToTimeFormat(codebase.timestamp)

		_, err = stmt.Exec(
			codebase.hash,
			codebase.sourcePath,
			codebase.language,
			codebase.cpgPath,
			jsonMetadata,
			timestampFormatted,
			timestampFormatted,
		)
		if err != nil {
			log.Fatalf("Couldn't insert: %#v\n", codebase)
		}
	}
}

func FindAllCodebases(db *sql.DB) ([]Codebase, error) {
	rows, err := db.Query(SelectAllStatement)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if cErr := rows.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}(rows)

	var total []Codebase
	for rows.Next() {
		var codebase Codebase
		var metadataJson []byte
		var timestamp *string

		err = rows.Scan(
			&codebase.hash,
			&codebase.sourcePath,
			&codebase.language,
			&codebase.cpgPath,
			&metadataJson,
			&timestamp,
		)
		if err != nil {
			return nil, err
		}

		if metadataJson != nil {
			var metadata Metadata
			err := json.Unmarshal(metadataJson, &metadata)
			if err != nil {
				return nil, err
			}

			codebase.metadata = metadata
		}
		if timestamp != nil {
			parsed, err := time.Parse(MainTimeFormat, *timestamp)
			if err != nil {
				return nil, err
			}
			codebase.timestamp = parsed
		}
		total = append(total, codebase)
	}

	return total, nil
}

func CreateCodebases(config Config, repositories []Repository, timestamp time.Time) []Codebase {
	total := make([]Codebase, len(repositories))
	for i, repository := range repositories {
		metadata := Metadata{
			ContainerCodebasePath: path.Join(config.JoernCodebasesAbsoluteBasePath, repository.hash),
			ContainerCpgPath:      path.Join(config.JoernCpgAbsoluteBasePath, repository.hash, CpgBinName),
			Status:                DefaultStatus,
		}
		codebase := Codebase{
			hash:       repository.hash,
			sourcePath: path.Join(config.CodebadgerCodebaseAbsoluteBasePath, repository.hash),
			language:   repository.language,
			cpgPath:    path.Join(config.CodebadgerCpgBasePath, repository.hash, CpgBinName),
			metadata:   metadata,
			timestamp:  timestamp,
		}
		total[i] = codebase
	}
	return total
}

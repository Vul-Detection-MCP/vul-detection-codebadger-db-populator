package main

import (
	"slices"
	"testing"
	"time"
)

func TestInsertCodebases(t *testing.T) {
	given := []Codebase{
		{
			hash:       "hash",
			sourcePath: "sourcePath",
			language:   "language",
			cpgPath:    "cpgPath",
			metadata: Metadata{
				ContainerCodebasePath: "ContainerCodebasePath",
				ContainerCpgPath:      "ContainerCpgPath",
				Status:                "Status",
			},
			timestamp: GetStableTime(),
		},
		{
			hash:       "hash2",
			sourcePath: "sourcePath2",
			language:   "language2",
			cpgPath:    "cpgPath2",
			metadata: Metadata{
				ContainerCodebasePath: "ContainerCodebasePath2",
				ContainerCpgPath:      "ContainerCpgPath2",
				Status:                "Status2",
			},
			timestamp: GetStableTime(),
		},
	}
	db := OpenDb("sqlite", ":memory:")

	ApplyBaseMigrations(db)
	InsertCodebases(db, given)
	actual, err := FindAllCodebases(db)
	if err != nil {
		panic(err)
	}

	expected := []Codebase{
		{
			hash:       "hash",
			sourcePath: "sourcePath",
			language:   "language",
			cpgPath:    "cpgPath",
			metadata: Metadata{
				ContainerCodebasePath: "ContainerCodebasePath",
				ContainerCpgPath:      "ContainerCpgPath",
				Status:                "Status",
			},
			timestamp: GetStableTime(),
		},
		{
			hash:       "hash2",
			sourcePath: "sourcePath2",
			language:   "language2",
			cpgPath:    "cpgPath2",
			metadata: Metadata{
				ContainerCodebasePath: "ContainerCodebasePath2",
				ContainerCpgPath:      "ContainerCpgPath2",
				Status:                "Status2",
			},
			timestamp: GetStableTime(),
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("Got: %#v, want: %#v", actual, expected)
	}
}

func TestCreateCodebases(t *testing.T) {
	repositories := []Repository{
		{
			hash:     "hash1",
			language: "c",
		},
		{
			hash:     "hash2",
			language: "cpp",
		},
	}
	configuration := Config{
		CodebadgerCodebaseAbsoluteBasePath: "/app/git",
		CodebadgerCpgBasePath:              "/app/playground/cpgs",
		JoernCodebasesAbsoluteBasePath:     "/playground/codebases",
		JoernCpgAbsoluteBasePath:           "/playground/cpgs",
	}
	timestamp := time.Now()

	actual := CreateCodebases(configuration, repositories, timestamp)

	expected := []Codebase{
		{
			hash:       "hash1",
			sourcePath: "/app/git/hash1",
			language:   "c",
			cpgPath:    "/app/playground/cpgs/hash1/cpg.bin",
			metadata: Metadata{
				ContainerCodebasePath: "/playground/codebases/hash1",
				ContainerCpgPath:      "/playground/cpgs/hash1/cpg.bin",
				Status:                "ready",
			},
			timestamp: timestamp,
		},
		{
			hash:       "hash2",
			sourcePath: "/app/git/hash2",
			language:   "cpp",
			cpgPath:    "/app/playground/cpgs/hash2/cpg.bin",
			metadata: Metadata{
				ContainerCodebasePath: "/playground/codebases/hash2",
				ContainerCpgPath:      "/playground/cpgs/hash2/cpg.bin",
				Status:                "ready",
			},
			timestamp: timestamp,
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("Got: %#v, want: %#v", actual, expected)
	}
}

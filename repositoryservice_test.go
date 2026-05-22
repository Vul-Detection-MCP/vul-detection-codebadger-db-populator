package main

import (
	"slices"
	"strings"
	"testing"
)

func TestParseCsv2(t *testing.T) {
	given := `hash1,c
hash2,cpp
hash3,c`

	a := strings.NewReader(given)
	actual := ParseCsvToRepositories(a)

	expected := []Repository{
		{
			hash:     "hash1",
			language: "c",
		},
		{
			hash:     "hash2",
			language: "cpp",
		},
		{
			hash:     "hash3",
			language: "c",
		},
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("Got: %#v, want: %#v", actual, expected)
	}
}

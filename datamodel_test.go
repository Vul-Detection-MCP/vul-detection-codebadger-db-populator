package main

import (
	"encoding/json"
	"testing"
)

func TestMetadataSerialization(t *testing.T) {
	given := Metadata{
		ContainerCodebasePath: "codebase path",
		ContainerCpgPath:      "cpg path",
		Status:                "active",
	}

	bytes, _ := json.Marshal(given)
	actual := string(bytes)

	expected := "{\"container_codebase_path\":\"codebase path\",\"container_cpg_path\":\"cpg path\",\"repository\":null,\"status\":\"active\"}"

	if actual != expected {
		t.Errorf("Got: %s, want: %s", actual, expected)
	}
}

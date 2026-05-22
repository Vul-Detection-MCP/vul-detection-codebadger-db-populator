package main

import "time"

// Repository is to get from CSV
type Repository struct {
	hash     string
	language string
}

type Codebase struct {
	hash       string
	sourcePath string
	language   string
	cpgPath    string
	metadata   Metadata
	timestamp  time.Time
}

type Metadata struct {
	ContainerCodebasePath string  `json:"container_codebase_path"`
	ContainerCpgPath      string  `json:"container_cpg_path"`
	Repository            *string `json:"repository"`
	Status                string  `json:"status"`
}

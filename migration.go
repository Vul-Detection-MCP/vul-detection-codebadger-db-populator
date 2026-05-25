package main

import "database/sql"

func createCodebasesTable(db *sql.DB) {
	_, _ = db.Exec(`
CREATE TABLE codebases (
	hash TEXT PRIMARY KEY,
	source_type TEXT NOT NULL DEFAULT 'local',
	source_path TEXT NOT NULL,
	language TEXT NOT NULL,
	cpg_path TEXT NOT NULL,
	joern_port INTEGER NOT NULL,
	metadata TEXT NOT NULL,
	created_at TEXT NOT NULL,
	last_accessed TEXT NOT NULL
);`)
}

func createFindingsTable(db *sql.DB) {
	_, _ = db.Exec(`
CREATE TABLE findings (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	codebase_hash TEXT NOT NULL,
	finding_type TEXT NOT NULL,
	severity TEXT NOT NULL,
	confidence TEXT NOT NULL,
	filename TEXT NOT NULL,
	line_number INTEGER NOT NULL,
	message TEXT NOT NULL,
	description TEXT,
	cwe_id INTEGER,
	rule_id TEXT,
	flow_data TEXT,
	metadata TEXT,
	created_at TEXT NOT NULL,
	FOREIGN KEY (codebase_hash) REFERENCES codebases(hash)
);`)
	_, _ = db.Exec("CREATE INDEX idx_findings_codebase ON findings(codebase_hash)")
	_, _ = db.Exec("CREATE INDEX idx_findings_severity ON findings(severity)")
	_, _ = db.Exec("CREATE INDEX idx_findings_confidence ON findings(confidence)")
	_, _ = db.Exec("CREATE INDEX idx_findings_type ON findings(finding_type)")
}

func createToolCacheTable(db *sql.DB) {
	_, _ = db.Exec(`
CREATE TABLE tool_cache (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	tool_name TEXT,
	codebase_hash TEXT,
	parameters_hash TEXT,
	parameters TEXT,
	output TEXT,
	created_at TEXT,
	UNIQUE(tool_name, codebase_hash, parameters_hash)
);`)
}

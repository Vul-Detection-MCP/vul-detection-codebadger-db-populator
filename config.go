package main

import "os"

type Config struct {
	DbPath                             string
	RepositoriesCsvAbsolutePath        string
	CodebadgerCodebaseAbsoluteBasePath string
	CodebadgerCpgBasePath              string
	JoernCodebasesAbsoluteBasePath     string
	JoernCpgAbsoluteBasePath           string
	JoernMinPort                       string
	JoernMaxPort                       string
}

var config Config

func envOrPanic(k string) string {
	v, present := os.LookupEnv(k)
	if !present {
		panic(k + "variable is not set.")
	}
	return v
}

func GetConfig() Config {
	return config
}

func InitConfig() {
	config = Config{
		DbPath:                             envOrPanic("DB_ABSOLUTE_PATH"),
		RepositoriesCsvAbsolutePath:        envOrPanic("CSV_ABSOLUTE_PATH"),
		CodebadgerCodebaseAbsoluteBasePath: envOrPanic("DB_RELATED_CODEBADGER_CODEBASES_ABSOLUTE_BASE_PATH"),
		CodebadgerCpgBasePath:              envOrPanic("DB_RELATED_CODEBADGER_CPG_ABSOLUTE_BASE_PATH"),
		JoernCodebasesAbsoluteBasePath:     envOrPanic("DB_RELATED_JOERN_CODEBASES_ABSOLUTE_BASE_PATH"),
		JoernCpgAbsoluteBasePath:           envOrPanic("DB_RELATED_JOERN_CPG_ABSOLUTE_BASE_PATH"),
		JoernMinPort:                       envOrPanic("DB_RELATED_JOERN_MIN_PORT"),
		JoernMaxPort:                       envOrPanic("DB_RELATED_JOERN_MAX_PORT"),
	}
}

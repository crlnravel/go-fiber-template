package config

import (
	"os"

	"github.com/joho/godotenv"
)

var LOADED = false

const (
	EnvironmentProduction = iota
	EnvironmentStaging
	EnvironmentDevelopment
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("env is not loaded properly")
	}
}

func GetEnv(key, fallback string) string {
	if !LOADED {
		loadEnv()
		LOADED = true
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetStageStatus() int {
	switch GetEnv("STAGE_STATUS", "dev") {
	case "prod":
		return EnvironmentProduction
	case "staging":
		return EnvironmentStaging
	case "dev":
		return EnvironmentDevelopment
	}

	return -1
}

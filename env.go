package main

import (
	"fmt"
	"os"
	"strconv"
)

func FetchEnvBool(key string) bool {
	value := FetchEnvStringWithDefault(key, "false")
	valueParsed, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return valueParsed
}

func FetchEnvString(key string) string {
	return FetchEnvStringWithDefault(key, "")
}

func EnvMustString(key string) (string, error) {
	value := FetchEnvStringWithDefault(key, "")
	if len(value) == 0 {
		return "", fmt.Errorf("the value for %s key is blank", key)
	}

	return value, nil
}

func FetchEnvStringWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

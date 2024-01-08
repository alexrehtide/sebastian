package configservice

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getDurationSecondsEnv(key string, defaultVal time.Duration) (time.Duration, error) {
	value := getEnv(key, "")
	if value == "" {
		return defaultVal, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return time.Duration(int64(intValue)), nil
}

func getBoolEnv(key string, defaultVal bool) (bool, error) {
	value := getEnv(key, "false")
	value = strings.ToLower(value)
	if value == "true" {
		return true, nil
	} else if value == "false" {
		return false, nil
	}
	return false, fmt.Errorf("configservice.Service.Load: ENV %s must be 'true' or 'false'", key)
}

func getIntEnv(key string, defaultVal int) (int, error) {
	value := getEnv(key, "")
	if value == "" {
		return defaultVal, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return intValue, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

package configs

import (
	"strconv"
	"strings"
	"time"
)

type EnvConfig map[string]string

func (config EnvConfig) GetOrDefault(key string, defaultValue string) string {
	if value, isExist := config[key]; isExist {
		return value
	}

	return defaultValue
}

func (config EnvConfig) GetOrDefaultDuration(
	key string,
	defaultValue time.Duration,
) (time.Duration, error) {
	if value, isExist := config[key]; isExist {
		return time.ParseDuration(value)
	}

	return defaultValue, nil
}

func (config EnvConfig) GetOrDefaultInt(
	key string,
	defaultValue int,
) (int, error) {
	value, isExist := config[key]
	if !isExist {
		return defaultValue, nil
	}

	trimmedValue := strings.TrimSpace(value)

	number, err := strconv.Atoi(trimmedValue)
	if err != nil {
		return 0, err
	}

	return number, nil
}

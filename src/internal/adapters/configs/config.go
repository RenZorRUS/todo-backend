package configs

type EnvConfig map[string]string

func (config EnvConfig) GetOrDefault(key string, defaultValue string) string {
	if value, isExist := config[key]; isExist {
		return value
	}

	return defaultValue
}

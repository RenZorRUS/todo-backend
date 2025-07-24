package configs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvConfig_GetOrDefault_GetExistValue(t *testing.T) {
	t.Parallel()

	config := EnvConfig{"key": "value"}
	value := config.GetOrDefault("key", "default")

	require.Equal(t, "value", value)
}

func TestEnvConfig_GetOrDefault_GetDefaultValue(t *testing.T) {
	t.Parallel()

	config := EnvConfig{"key": "value"}
	value := config.GetOrDefault("key2", "default")

	require.Equal(t, "default", value)
}

package utils

import (
	"fmt"
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAppMode_Success(t *testing.T) {
	cases := []tests.TestCase[string, string]{
		{
			Name: fmt.Sprintf(
				"Should return app mode if %s env variable is specified",
				consts.AppMode,
			),
			Input:    consts.ProdMode,
			Expected: consts.ProdMode,
		},
		{
			Name: fmt.Sprintf(
				"Should return %s if %s env variable is not specified",
				consts.DevMode,
				consts.AppMode,
			),
			Input:    "",
			Expected: consts.DevMode,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			if test.Input != "" {
				t.Setenv(consts.AppMode, test.Input)
			}

			result := GetAppMode()

			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestIsProdMode(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[string, bool]{
		{
			Name:     consts.AppMode + " empty value, should return false",
			Input:    "",
			Expected: false,
		},
		{
			// Same as empty – `os.LookupEnv` returns "", false
			Name:     consts.AppMode + " missing variable. should return false",
			Input:    "",
			Expected: false,
		},
		{
			Name:     consts.AppMode + " dev lowercase, should return false",
			Input:    "dev",
			Expected: false,
		},
		{
			Name:     consts.AppMode + " DEV uppercase, should return false",
			Input:    "DEV",
			Expected: false,
		},
		{
			Name:     consts.AppMode + " prod lowercase, should return true",
			Input:    "prod",
			Expected: true,
		},
		{
			Name:     consts.AppMode + " PROD uppercase, should return true",
			Input:    "PROD",
			Expected: true,
		},
		{
			Name:     consts.AppMode + " production lowercase, should return true",
			Input:    "production",
			Expected: true,
		},
		{
			Name:     consts.AppMode + " PRODUCTION uppercase, should return true",
			Input:    "PRODUCTION",
			Expected: true,
		},
		{
			Name:     consts.AppMode + " garbage text, should return false",
			Input:    "staging",
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result := IsProdMode(test.Input)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestGetEnvFilePath_Success(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[string, string]{
		{
			Name:     "If app in prod mode, func should return " + consts.ProdEnvFile,
			Input:    consts.ProdMode,
			Expected: consts.ProdEnvFile,
		},
		{
			Name:     "If app in dev mode, func should return " + consts.DevEnvFile,
			Input:    consts.DevMode,
			Expected: consts.DevEnvFile,
		},
		{
			Name:     "If app in test mode, func should return " + consts.TestEnvFile,
			Input:    consts.TestMode,
			Expected: consts.TestEnvFile,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, err := GetEnvFilePath(test.Input)

			require.NoError(t, err)
			assert.Equal(t, test.Expected, result)
		})
	}
}

func TestGetEnvFilePath_Fail(t *testing.T) {
	t.Parallel()

	result, err := GetEnvFilePath("unknown")

	require.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, errs.ErrUnknownAppMode, err)
}

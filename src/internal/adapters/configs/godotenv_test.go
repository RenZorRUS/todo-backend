package configs

import (
	"fmt"
	"io/fs"
	"syscall"
	"testing"

	"github.com/RenZorRUS/todo-backend/src/internal/adapters/consts"
	"github.com/RenZorRUS/todo-backend/src/internal/adapters/errs"
	"github.com/RenZorRUS/todo-backend/src/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGoDotEnvLoader(t *testing.T) {
	t.Parallel()

	loader := NewGoDotEnvLoader()

	require.NotNil(t, loader)
}

// Its not parallel because it creates file that fails other loader tests.
func TestGoDotEnvLoader_Load_Success(t *testing.T) { //nolint:paralleltest
	fileUtils := tests.NewTestFileUtils(t)
	loader := NewGoDotEnvLoader()

	currentDir := fileUtils.GetCurrentDir()
	filePath := fmt.Sprintf("%s/%s", currentDir, consts.TestEnvFile)

	fileUtils.CreateFile(filePath)
	defer fileUtils.RemoveFile(filePath)

	appConfig, err := loader.Load(consts.TestMode)

	require.NoError(t, err)
	require.NotNil(t, appConfig)
}

func TestGoDotEnvLoader_Load_Fail(t *testing.T) {
	t.Parallel()

	cases := []tests.TestCase[string, any]{
		{
			Name:     "Should return error if app mode is unknown",
			Input:    "unknown",
			Expected: errs.ErrUnknownAppMode,
		},
		{
			Name:  "Should return error if test env file is not found",
			Input: consts.TestMode,
			Expected: &fs.PathError{
				Op:   "open",
				Path: consts.TestEnvFile,
				Err:  syscall.Errno(2),
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			loader := NewGoDotEnvLoader()

			appConfig, err := loader.Load(test.Input)

			require.Error(t, err)
			assert.Nil(t, appConfig)
			assert.Equal(t, test.Expected, err)
		})
	}
}

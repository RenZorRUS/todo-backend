package tests

import (
	"os"
	"path/filepath"
	"testing"
)

type TestFileUtils struct {
	t *testing.T
}

func NewTestFileUtils(t *testing.T) *TestFileUtils {
	t.Helper()
	return &TestFileUtils{t: t}
}

func (tfu *TestFileUtils) GetCurrentDir() string {
	tfu.t.Helper()

	currentDir, err := os.Getwd()
	if err != nil {
		tfu.t.Fatal(err)
	}

	return currentDir
}

func (tfu *TestFileUtils) CreateFile(filePath string) *os.File {
	tfu.t.Helper()

	filePath = filepath.Clean(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		tfu.t.Fatal(err)
	}

	return file
}

func (tfu *TestFileUtils) RemoveFile(filePath string) {
	tfu.t.Helper()

	err := os.Remove(filePath)
	if err != nil {
		tfu.t.Fatal(err)
	}
}

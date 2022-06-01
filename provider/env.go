package env

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rajasoun/aws-hub/test"
)

// Get Working directory function
// To Enable Mocking - For TDD Error Handling coverage
func GetWorkingDirectory(osGetwd func() (string, error)) (string, error) {
	dirpath, err := osGetwd()
	if err != nil {
		return "", errors.New("could not get current working directory")
	}
	dirpath = correctPathIfExecutedFromTest(dirpath)
	return dirpath, nil
}

func correctPathIfExecutedFromTest(dirpath string) string {
	if test.IsTestRun() {
		dirpath = removeLeafDir(dirpath)
	}
	return dirpath
}

func removeLeafDir(dirpath string) string {
	leafDir := filepath.Base(dirpath)
	return strings.ReplaceAll(dirpath, leafDir, "")
}

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

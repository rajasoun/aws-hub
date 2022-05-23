package providers

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWorkingDirectory(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check GetWorkingDirectory", func(t *testing.T) {
		tests := []struct {
			testName    string
			notContains string
			want        bool
		}{
			{"Check aws-hub in dirPath should Pass", "aws-hub", true},
			{"Check providers in dirPath should Fail", "providers", false},
		}
		dirpath, _ := GetWorkingDirectory(os.Getwd)
		for _, test := range tests {
			t.Run(test.testName, func(t *testing.T) {
				got := strings.Contains(dirpath, test.notContains)
				assert.Equal(test.want, got,
					"GetWorkingDirectory(os.Getwd) = %v,\nCondition: Not Contains %v, got = %v want = %v",
					dirpath, test.notContains, got, test.want)
			})
		}
	})
	t.Run("Check GetWorkingDirectory Handles Error Gracefully", func(t *testing.T) {
		wantErrMsg := "could not get current working directory"
		// Mocked function os.Getwd
		osGetwd := func() (string, error) {
			err := errors.New(wantErrMsg)
			return "", err
		}
		// This will return error
		_, err := GetWorkingDirectory(osGetwd)
		gotErrMsg := err.Error()
		assert.Equal(gotErrMsg, wantErrMsg, "GetWorkingDirectory() = %v , want = %v",
			err, wantErrMsg)
	})
}

func TestFileExists(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		testName string
		filePath string
		want     bool
	}{
		{"Check main.go Exists in Base Dir", "main.go", true},
		{"Check dummy.go Does Not Exists in Base Dir", "dummy.go", false},
	}
	basePath, _ := GetWorkingDirectory(os.Getwd)
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			filePath := basePath + test.filePath
			got := FileExists(filePath)
			assert.Equal(test.want, got, "FileExists(%v) = %v , want = %v", filePath, got, test.want)
		})
	}
}

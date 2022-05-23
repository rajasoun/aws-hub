package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTestRun(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check If from Test Run", func(t *testing.T) {
		want := true
		got := IsTestRun()
		assert.Equal(want, got, "IsTestRun() = %v, want %v", got, want)

	})
}

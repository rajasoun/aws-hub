package credential

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefaultProfile(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check Loading Default Profile with Region", func(t *testing.T) {
		cfg, err := LoadDefaultProfile()
		got := cfg.Region
		want := "us-east-1"
		assert.NoError(err, "LoadDefaultProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
}

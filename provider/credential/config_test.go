package credential

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefaultProfile(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	credLoader := new(CredentialLoader)
	t.Run("Check Load Default Profile with Region", func(t *testing.T) {
		cfg, err := credLoader.LoadDefaultProfile()
		got := cfg.Region
		want := "us-east-1"
		assert.NoError(err, "LoadDefaultProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
	t.Run("Check Load Default Profile For Err", func(t *testing.T) {
		_, err := credLoader.LoadDefaultProfile()
		assert.Error(err, "LoadDefaultProfile() error = %v", err)
	})
}

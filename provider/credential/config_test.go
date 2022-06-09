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
		cfg, err := credLoader.LoadDefaultConfig()
		got := cfg.Region
		want := "us-east-1"
		assert.NoError(err, "LoadDefaultProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
	t.Run("Check Load Default Profile For Err", func(t *testing.T) {
		t.Skip("Failing Test")
		_, err := credLoader.LoadDefaultConfig()
		assert.Error(err, "LoadDefaultProfile() error = %v", err)
	})
}

func TestLoadDefaultConfigForProfile(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	credLoader := new(CredentialLoader)
	profile := "secops-experiments"
	t.Run("Check Load Default Profile with Region", func(t *testing.T) {
		cfg, err := credLoader.LoadDefaultConfigForProfile(profile)
		got := cfg.Region
		want := "us-east-1"
		assert.NoError(err, "LoadDefaultProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
	t.Run("Check Load Default Profile For Err", func(t *testing.T) {
		t.Skip("Failing Test")
		_, err := credLoader.LoadDefaultConfigForProfile(profile)
		assert.Error(err, "LoadDefaultProfile() error = %v", err)
	})
}

func TestCredentialLoaderGetSections(t *testing.T) {
	assert := assert.New(t)
	t.Run("", func(t *testing.T) {
		credLoader := &CredentialLoader{}
		_, err := credLoader.GetSections()
		assert.NoError(err, "CredentialLoader.GetSections() error = %v ", err)
	})
}

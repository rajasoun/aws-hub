package credential

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestLoadCredentialFromFileForProfile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name    string
		profile string
		want    aws.Config
		//wantErr bool
	}{
		{
			name:    "Checl Loading Default Profile",
			profile: "Default",
			want: aws.Config{
				Region: DefaultRegion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.profile)
			assert.NoError(err, "LoadCredentialFromFileForProfile() error = %v", err)
			assert.Equal(got.Region, tt.want.Region, "LoadCredentialFromFileForProfile() = %v, want %v", got, tt.want)
		})
	}
}

func TestCheckConfig(t *testing.T) {
	if os.Getenv("SKIP_E2E") == "" {
		t.Skip("Skip INTEGRATION Tests")
	}
	assert := assert.New(t)
	t.Parallel()
	t.Run("Check Config for Valid Credentials", func(t *testing.T) {
		cfg, err := LoadConfig("default")
		assert.NoError(err, "LoadConfig() error = %v", err)
		assert.NoError(err)
		result := CheckConfig(cfg)
		assert.True(result, "Check LoadConfig() = %v", result)
	})
}

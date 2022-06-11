package credential

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
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
		assert.NoError(err, "LoadDefaultConfigForProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
}

func TestCredentialLoaderGetSections(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check GetSections if Credential File Exists",
			args: args{
				filename: config.DefaultSharedCredentialsFilename(),
			},
			want: true,
		},
		{
			name: "Check GetSections if Credential File Not Exists",
			args: args{
				filename: ".aws/credentials",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			credLoader := &CredentialLoader{}
			got, _ := credLoader.GetSections(tt.args.filename)
			want := 0
			assert.GreaterOrEqual(len(got.List()), want,
				"CredentialLoader.GetSections() = %v , want = %v", len(got.List()), want)
		})
	}
}

func TestFileExists(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	baseDir, _ := os.Getwd()
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check File with Relative Path",
			args: args{
				filename: ".aws/credentials",
			},
			want: false,
		},
		{
			name: "Check File with Relative Path",
			args: args{
				filename: baseDir + "/config.go",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fileExists(tt.args.filename)
			assert.Equal(tt.want, got, "fileExists() = %v, want = %v for file %v", got, tt.want, tt.args.filename)
		})
	}
}

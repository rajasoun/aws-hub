package credential

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadDefaultProfile(t *testing.T) {
	testProfile := "secops-experiments"
	testRegion := "us-east-1"
	assert := assert.New(t)
	t.Parallel()
	credLoader := New()
	t.Run("Check Load Default Profile with Region", func(t *testing.T) {
		credLoader.LocalLoaderFunc = config.LoadDefaultConfig
		cfg, err := credLoader.LoadDefaultConfig()
		got := cfg.Region
		want := testRegion
		assert.NoError(err, "LoadDefaultProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
	t.Run("Check Load Default Profile with Region with Err", func(t *testing.T) {
		credLoader.LocalLoaderFunc = func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error) {
			return aws.Config{}, errors.New("simulated error")
		}
		_, err := credLoader.LoadDefaultConfig()
		assert.Error(err, "LoadDefaultConfig() error = %v", err)
	})
	t.Run("Check Load Default Profile with Region", func(t *testing.T) {
		credLoader.LocalLoaderFunc = config.LoadDefaultConfig
		cfg, err := credLoader.LoadDefaultConfigForProfile(testProfile)
		got := cfg.Region
		want := testRegion
		assert.NoError(err, "LoadDefaultConfigForProfile() error = %v", err)
		assert.Equal(got, want, "LoadDefaultProfile() = %v, want %v", got, want)
	})
	t.Run("Check Load Default Profile with Region", func(t *testing.T) {
		credLoader.LocalLoaderFunc = func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error) {
			return aws.Config{}, errors.New("simulated error")
		}
		_, err := credLoader.LoadDefaultConfigForProfile(testProfile)
		assert.Error(err, "LoadDefaultConfigForProfile() error = %v", err)
	})
}

func TestCredentialLoaderGetSections(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		filename string
	}
	tests := []struct {
		name            string
		args            args
		emptyCredential bool
	}{
		{
			name: "Check GetSections if Credential File Exists",
			args: args{
				filename: config.DefaultSharedCredentialsFilename(),
			},
		},
		{
			name: "Check GetSections if Credential File Not Exists",
			args: args{
				filename: ".aws/credentials",
			},
		},
		{
			name: "Check GetSections with Custom Location with Empty File",
			args: args{
				filename: "/tmp/credentials",
			},
			emptyCredential: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.emptyCredential {
				os.Create("/tmp/credentials")
			}
			credLoader := &Loader{}
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

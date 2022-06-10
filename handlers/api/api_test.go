package api

import (
	"bytes"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAwsAPI(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		apiName string
	}{
		{
			name:    "Check NewAwsAPI for GetUserCountAPI ",
			apiName: IAMGetUserCountAPI,
		},
		{
			name:    "Check NewAwsAPI for GetUserIdentityAPI",
			apiName: IAMGetUserIdentityAPI,
		},
		{
			name:    "Check NewAwsAPI for GetAliasesAPI",
			apiName: IAMGetAliasesAPI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewAwsAPI(tt.apiName)
			got := reflect.TypeOf(api).Name()
			assert.Equal(got, tt.apiName, "NewAwsAPI() = %v, want = %v ", got, tt.apiName)
		})
	}
}

func TestGetConfig(t *testing.T) {
	funcName := "GetConfig() = %v "
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name              string
		profile           string
		isMultipleProfile bool
		wantErr           bool
	}{
		{
			name:              "Check GetConfig for Empty Profile & isMultipleProfile is false",
			profile:           "",
			isMultipleProfile: false,
			wantErr:           false,
		},
		{
			name:              "Check GetConfig for invalid Profile & isMultipleProfile is false",
			profile:           "invalid-profile",
			isMultipleProfile: true,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := GetConfigFromFileSystem(tt.profile, tt.isMultipleProfile)
			if tt.wantErr {
				assert.Error(err, funcName, err)
			}
			assert.NoError(err, funcName, err)
			assert.NotEmpty(cfg.Region, funcName, cfg.Region)
		})
	}
}

func TestHandleErr(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "Check Log With  No Err",
			err:  nil,
			want: "successfuly",
		},
		{
			name: "Check Log With  Err",
			err:  errors.New("simulated error"),
			want: "Failed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outputBuffer bytes.Buffer
			log.SetOutput(&outputBuffer)
			log.SetFlags(0)
			handleErr(tt.err, "Test")
			got := outputBuffer.String()
			assert.Contains(got, tt.want, "handleErr() = %v, want = %v", got, tt.want)
		})
	}
}

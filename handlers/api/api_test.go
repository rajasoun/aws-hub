package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAwsAPI(t *testing.T) {
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
			assert := assert.New(t)
			api := NewAwsAPI(tt.apiName)
			got := reflect.TypeOf(api).Name()
			assert.Equal(got, tt.apiName, "NewAwsAPI() = %v, want = %v ", got, tt.apiName)
		})
	}
}

func TestGetConfig(t *testing.T) {
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
			assert := assert.New(t)
			cfg, err := GetConfigFromFileSystem(tt.profile, tt.isMultipleProfile)
			if tt.wantErr {
				assert.Error(err, "GetConfig() = %v ", err)
			}
			assert.NoError(err, "GetConfig() = %v ", err)
			assert.NotEmpty(cfg.Region, "GetConfig() = %v ", cfg.Region)
		})
	}
}

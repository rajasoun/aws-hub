package aws

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		apiName string
	}{
		{
			name:    "Check New for GetUserCountAPI ",
			apiName: IAMGetUserCountAPI,
		},
		{
			name:    "Check New for GetUserIdentityAPI",
			apiName: IAMGetUserIdentityAPI,
		},
		{
			name:    "Check New for GetAliasesAPI",
			apiName: IAMGetAliasesAPI,
		},
		{
			name:    "Check New for Ping",
			apiName: IAMPing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := New(tt.apiName)
			got := reflect.TypeOf(api).Name()
			assert.Equal(got, tt.apiName, "New() = %v, want = %v ", got, tt.apiName)
		})
	}
}

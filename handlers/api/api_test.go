package api

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAWSAPI(t *testing.T) {
	type args struct {
		api               string
		profile           string
		isMultipleProfile bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Check GetUserCountAPI with Empty Profile",
			args: args{
				api:               IAMGetUserCountAPI,
				profile:           "",
				isMultipleProfile: false,
			},
		},
		{
			name: "Check GetUserIdentityAPI with Empty Profile",
			args: args{
				api:               IAMGetUserIdentityAPI,
				profile:           "",
				isMultipleProfile: false,
			},
		},
		{
			name: "Check GetAliasesAPI with Empty Profile",
			args: args{
				api:               IAMGetAliasesAPI,
				profile:           "",
				isMultipleProfile: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			_, err := GetConfig(tt.args.profile, tt.args.isMultipleProfile)
			assert.NoError(err, "GetConfig() Err = %v ", err)
			api := NewAwsAPI(tt.args.api)
			got := reflect.TypeOf(api).Name()
			assert.Equal(got, tt.args.api, "NewAwsAPI() = %v, want = %v ", got, tt.args.api)
			// client := iam.NewFromConfig(aws.Config{})
			// api := NewAwsAPI(tt.args.api)
			// got, err := api.Execute(client)
			// assert.Error(err, "Execute() Err = %v ", err)
			// if tt.wantEmpty {
			// 	assert.Empty(got, "Execute() = %v", got)
			// }
		})
	}
}

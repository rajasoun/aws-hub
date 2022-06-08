package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAwsAPI(t *testing.T) {
	if os.Getenv("SKIP_E2E") != "" {
		t.Skip("Skipping INTEGRATION Tests")
	}
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
				api:               "GetUserCount",
				profile:           "",
				isMultipleProfile: false,
			},
		},
		{
			name: "Check GetUserIdentityAPI with Empty Profile",
			args: args{
				api:               "GetUserIdentity",
				profile:           "",
				isMultipleProfile: false,
			},
		},
		{
			name: "Check GetAliasesAPI with Empty Profile",
			args: args{
				api:               "GetAliases",
				profile:           "",
				isMultipleProfile: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			api := NewAwsAPI(tt.args.api)
			cfg, err := GetConfig(tt.args.profile, tt.args.isMultipleProfile)
			assert.NoError(err, "GetConfig() Err = %v ", err)
			got, err := api.Execute(cfg)
			assert.NoError(err, "Execute() Err = %v ", err)
			assert.NotEmpty(got, "Execute() = %v, want %v", got)
		})
	}
}

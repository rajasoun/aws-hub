package api

import (
	"testing"
)

func TestNewAwsAPI(t *testing.T) {
	// type args struct {
	// 	api               string
	// 	profile           string
	// 	isMultipleProfile bool
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	{
	// 		name: "Check GetUserCountAPI with Empty Profile",
	// 		args: args{
	// 			api:               "GetUserCount",
	// 			profile:           "",
	// 			isMultipleProfile: false,
	// 		},
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		assert := assert.New(t)
	// 		api := NewAwsAPI(tt.args.api)
	// 		cfg, err := GetConfig(tt.args.profile, tt.args.isMultipleProfile)
	// 		assert.NoError(err, "GetConfig() Err = %v ", err)
	// 		got, err := api.Execute(cfg)
	// 		assert.NoError(err, "Execute() Err = %v ", err)
	// 		assert.NotEmpty(got, "Execute() = %v, want %v", got)
	// 	})
	// }
}

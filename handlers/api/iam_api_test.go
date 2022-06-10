package api

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	type args struct {
		client *iam.Client
	}
	tests := []struct {
		name    string
		api     AwsAPI
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Check GetAliasesAPI With NoOp Client",
			api:  NewAwsAPI(IAMGetAliasesAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserCountAPI With NoOp Client",
			api:  NewAwsAPI(IAMGetUserCountAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserIdentitytAPI With NoOp Client",
			api:  NewAwsAPI(IAMGetUserIdentityAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.api.Execute(tt.args.client)
			assert.Error(err, "GetAliasesAPI.Execute() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}

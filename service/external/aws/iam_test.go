package aws

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
		api     API
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Check GetAliasesAPI With NoOp Client",
			api:  New(IAMGetAliasesAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserCountAPI With NoOp Client",
			api:  New(IAMGetUserCountAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserIdentitytAPI With NoOp Client",
			api:  New(IAMGetUserIdentityAPI),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check Ping",
			api:  New(IAMPing),
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.api.Execute(tt.args.client)
			if tt.wantErr {
				assert.Error(err, "Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

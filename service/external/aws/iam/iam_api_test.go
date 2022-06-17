package iam

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
		name string
		api  interface {
			Execute(client *iam.Client) (interface{}, error)
		}
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Check GetAliasesAPI With NoOp Client",
			api:  GetAliasesAPI{},
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserCountAPI With NoOp Client",
			api:  GetUserCountAPI{},
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check GetUserIdentitytAPI With NoOp Client",
			api:  GetUserIdentityAPI{},
			args: args{
				client: &iam.Client{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check Ping",
			api:  DoPing{},
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

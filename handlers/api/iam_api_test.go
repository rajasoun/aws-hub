package api

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock AWS API
type MockAwsAPI struct {
	mock.Mock
}

type MockOutput struct {
	Message string
}

/**
* Mock using testify Framework
 */

// Mock Execute Function
func (c *MockAwsAPI) Execute(client *iam.Client) (interface{}, error) {
	args := c.Called(client)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*MockOutput), args.Error(1)
}

func TestGetAliasesAPI_Execute(t *testing.T) {
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

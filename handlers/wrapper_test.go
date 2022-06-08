package handlers

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/stretchr/testify/assert"
)

func TestAWSHandler_SdkWrapperAPI(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		apiName string
	}{
		{"Check IAMGetUserCountHandler", "GetUserCount"},
		{"Check IAMGetUserIdentityHandler", "GetUserIdentity"},
		{"Check IAMGetAliasesHandler", "GetAliases"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewDefaultAWSHandler(false)
			emptyCfg := aws.Config{}
			noOpClient := iam.NewFromConfig(emptyCfg)
			_, err := handler.SdkWrapperAPI(noOpClient, tt.apiName)
			assert.Error(err, "No Err Occured with Empty Profile. err = %v ", err)
		})
	}
}

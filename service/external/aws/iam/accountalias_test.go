package iam

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/**
* Mock using testify Framework
 */

// Mock Function to AWS ListAccountAliases
// Technique: Interface Substitution
// Interface Substitution is done by mocking methods that are implemented by an interface.
// Steps:
//	1. make an object of struct
//	2. implements all methods in the interface for mocking real implementation
func (mockClient *MockClient) ListAccountAliases(ctx context.Context,
	params *iam.ListAccountAliasesInput,
	optFns ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	// Mock ListAccountAliases of AWS
	// Mocked ListAccountAliases Function will be Called and Results Injected
	result := mockClient.Called(ctx, params, optFns)
	// On Error
	if result.Get(1) != nil {
		return result.Get(0).(*iam.ListAccountAliasesOutput), result.Error(1)
	}
	// If No Error
	return result.Get(0).(*iam.ListAccountAliasesOutput), nil
}

func TestListAccountAliasesViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name    string
		input   []string
		want    string
		wantErr error
	}{
		{
			name:    "Check ListAccountAliases",
			input:   []string{testAlias},
			want:    testAlias,
			wantErr: nil,
		},
		{
			name:    "Check ListAccountAliases with Err",
			input:   []string{},
			wantErr: errors.New("simulated error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := new(MockClient)
			expectedOutput := &iam.ListAccountAliasesOutput{
				AccountAliases: tt.input,
			}
			// Inject Mock Function to the Client
			client.
				On("ListAccountAliases", mock.Anything, mock.Anything, mock.Anything).
				Return(expectedOutput, tt.wantErr)
			got, err := GetAliases(client)
			if tt.wantErr != nil {
				assert.Error(err, "expect no error, got %v", err)
				assert.Empty(got.List, "GetAliases() = %v", got.List)
				return
			}
			assert.NoError(err, "GetAliases() %v", err)
			assert.Equal(testAlias, got.List[0], "GetAliases() = %v, want = %v", got.List[0], testAlias)
		})
	}
}

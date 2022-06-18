package cost

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/**
* Mock using testify Framework
 */

// Mock Function to AWS GetCostAndUsageAPI
// Technique: Interface Substitution
// Interface Substitution is done by mocking methods that are implemented by an interface.
// Steps:
//	1. make an object of struct
//	2. implements all methods in the interface for mocking real implementation
func (mockClient *MockClient) GetCostAndUsage(ctx context.Context, params *costexplorer.GetCostAndUsageInput,
	optFns ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageOutput, error) {
	// Mock GetCostAndUsage of AWS
	// Mocked GetCostAndUsage Function will be Called and Results Injected
	result := mockClient.Called(ctx, params, optFns)
	// On Error
	if result.Get(1) != nil {
		return result.Get(0).(*costexplorer.GetCostAndUsageOutput), result.Error(1)
	}
	// If No Error
	return result.Get(0).(*costexplorer.GetCostAndUsageOutput), nil
}

func TestGetCostAndUsageViaMockFramework(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name           string
		expectedOutput *costexplorer.GetCostAndUsageOutput
		want           float64
		wantErr        error
	}{
		{
			name:           "Check GetCostAndUsage",
			expectedOutput: &costexplorer.GetCostAndUsageOutput{},
			want:           0,
			wantErr:        nil,
		},
		{
			name:           "Check GetCostAndUsage with Err",
			expectedOutput: &costexplorer.GetCostAndUsageOutput{},
			want:           0,
			wantErr:        errors.New("simulated error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := new(MockClient)
			// Inject Mock Function to the Client
			client.
				On("GetCostAndUsage", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.expectedOutput, tt.wantErr)
			got, err := CurrentBill(client)
			if tt.wantErr != nil {
				assert.Error(err, "expect no error, got %v", err)
				assert.Equal(tt.want, got.Total, "GetAliases() = %v, want = %v", got.Total, tt.want)
				assert.Empty(got.Total, "GetAliases() = %v", got.Total)
				return
			}
			assert.Equal(tt.want, got.Total, "GetAliases() = %v, want = %v", got.Total, tt.want)
			assert.NoError(err, "GetAliases() %v", err)
		})
	}
}

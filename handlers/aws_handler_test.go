package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewDefaultAWSHandler(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name          string
		multiple      bool
		wantCacheType string
	}{
		{
			name:          "Check New Default AWS Handler with no profile",
			multiple:      false,
			wantCacheType: "InMemoryCache",
		},
		{
			name:          "Check New Default AWS Handler with multiple profile",
			multiple:      true,
			wantCacheType: "InMemoryCache",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultAWSHandler(tt.multiple)
			assert.Equal(got.cache.Type(), tt.wantCacheType, "NewDefaultAWSHandler() = %v, want %v", got, tt.wantCacheType)
			assert.Empty(got.GetAWSHandler(), "GetAWSHandler() = %v", got.GetAWSHandler())
		})
	}
}

// To mock AWS operations
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
func (mockFunc *MockAwsAPI) Execute(client *iam.Client) (interface{}, error) {
	// Call Mock Function call
	// The Function to be Called and Result will be Injected in Test
	result := mockFunc.Called(client)
	//Return Result On Error
	if result.Get(0) == nil {
		return nil, result.Error(1)
	}
	//Return Result If No Error
	return result.Get(0).(*MockOutput), result.Error(1)
}

func TestInvokeAPI(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	awsHandler := NewDefaultAWSHandler(false)
	mockServer := test.MockServer{}

	t.Run("Check InvokeAPI for Success", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/test", nil)
		responseWriter := mockServer.DoSimulation(test.MockSuccessHandler, nil)
		awsWrapper := AWSWrapper{
			request:  request,
			writer:   responseWriter,
			cache:    awsHandler.cache,
			multiple: awsHandler.multiple,
		}
		// Inject Mock Function to be Called along with Resturn values as Parameter
		client := new(MockAwsAPI)
		expectedOutput := &MockOutput{Message: "Test"}
		client.On("Execute", mock.Anything).Return(expectedOutput, nil)
		awsWrapper.InvokeAPI(client, "dummy", "dummy")
		//For Checking Cache
		awsWrapper.InvokeAPI(client, "dummy", "dummy")
	})

	t.Run("Check InvokeAPI for Err", func(t *testing.T) {
		responseWriter := mockServer.DoSimulation(test.MockFailureHandler, nil)
		gotStatusCode := responseWriter.Result().StatusCode
		assert.Equal(gotStatusCode, http.StatusInternalServerError)
		body, _ := io.ReadAll(responseWriter.Result().Body)
		gotBody := string(body)
		assert.Contains(gotBody, "error")
	})
}

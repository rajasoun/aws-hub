package handlers

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rajasoun/aws-hub/test"
	"github.com/stretchr/testify/mock"
)

// To mock AWS operations.
type MockAwsAPI struct {
	mock.Mock
}

type MockOutput struct {
	Message string
}

/**
* Mock using testify Framework
 */

// Mock Execute Function.
func (mockFunc *MockAwsAPI) Execute(client *iam.Client) (interface{}, error) {
	// Call Mock Function call
	// The Function to be Called and Result will be Injected in Test
	result := mockFunc.Called(client)
	// Return Result On Error
	if result.Get(0) == nil {
		return nil, result.Error(1)
	}
	// Return Result If No Error
	return result.Get(0).(*MockOutput), result.Error(1)
}

func TestInvokeAPI(t *testing.T) {
	t.Parallel()
	t.Run("Check InvokeAPI for Success", func(t *testing.T) {
		awsWrapper := MockAWSWrapper(test.MockSuccessHandler)
		expectedOutput := &MockOutput{Message: "Test with Success"}
		client := new(MockAwsAPI)
		client.On("Execute", mock.Anything).Return(expectedOutput, nil)
		awsWrapper.InvokeAPI(client, "dummy", "dummy")
		// Check Cache
		awsWrapper.InvokeAPI(client, "dummy", "dummy")
	})
	t.Run("Check InvokeAPI for Failure", func(t *testing.T) {
		awsWrapper := MockAWSWrapper(test.MockFailureHandler)
		expectedOutput := &MockOutput{Message: "Test with Failure"}
		client := new(MockAwsAPI)
		client.On("Execute", mock.Anything).Return(expectedOutput, errors.New("simulated error"))
		awsWrapper.InvokeAPI(client, "dummy", "Simulated Error")
	})
}

func MockAWSWrapper(handler func(w http.ResponseWriter, r *http.Request)) *AWSWrapper {
	awsHandler := NewDefaultAWSHandler(false)
	mockServer := test.MockServer{}
	request, _ := http.NewRequest("GET", "/test", nil)
	responseWriter := mockServer.DoSimulation(handler, nil)
	awsWrapper := AWSWrapper{
		request:  request,
		writer:   responseWriter,
		cache:    awsHandler.cache,
		multiple: awsHandler.multiple,
	}
	return &awsWrapper
}

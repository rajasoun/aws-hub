package handlers

import (
	"bytes"
	"errors"
	"log"
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
			assert.Equal(got.cache.Type(), tt.wantCacheType,
				"NewDefaultAWSHandler() = %v, want %v", got, tt.wantCacheType)
		})
	}
}

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

// Mock AWS Wrapper.
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

func TestInvokeAPI(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		awsWrapper func(responseWriter http.ResponseWriter, request *http.Request)
		msg        string
		wantErr    bool
	}{
		{
			name:       "Check InvokeAPI for Success",
			awsWrapper: test.MockSuccessHandler,
			msg:        "Test with Success",
			wantErr:    false,
		},
		{
			name:       "Check InvokeAPI for Failure",
			awsWrapper: test.MockFailureHandler,
			msg:        "Test with Failure",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			awsWrapper := MockAWSWrapper(tt.awsWrapper)
			expectedOutput := &MockOutput{Message: tt.msg}
			client := new(MockAwsAPI)
			if tt.wantErr {
				client.On("Execute", mock.Anything).Return(expectedOutput, errors.New("simulated error"))
				awsWrapper.InvokeAPI(client, "dummy", "Simulated Error")
				return
			}
			client.On("Execute", mock.Anything).Return(expectedOutput, nil)
			awsWrapper.InvokeAPI(client, "dummy", "dummy")
			// Check Cache
			awsWrapper.InvokeAPI(client, "dummy", "dummy")
		})
	}
}

func TestGetConfig(t *testing.T) {
	funcName := "GetConfig() = %v "
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name              string
		profile           string
		isMultipleProfile bool
		wantErr           bool
	}{
		{
			name:              "Check GetConfig for Empty Profile & isMultipleProfile is false",
			profile:           "",
			isMultipleProfile: false,
			wantErr:           false,
		},
		{
			name:              "Check GetConfig for invalid Profile & isMultipleProfile is false",
			profile:           "invalid-profile",
			isMultipleProfile: true,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := loadConfigFromFileSystem(tt.profile, tt.isMultipleProfile)
			if tt.wantErr {
				assert.Error(err, funcName, err)
			}
			assert.NoError(err, funcName, err)
			assert.NotEmpty(cfg.Region, funcName, cfg.Region)
		})
	}
}

func TestHandleErr(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "Check Log With  No Err",
			err:  nil,
			want: "successfully",
		},
		{
			name: "Check Log With  Err",
			err:  errors.New("simulated error"),
			want: "Failed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var outputBuffer bytes.Buffer
			log.SetOutput(&outputBuffer)
			log.SetFlags(0)
			handleErr(tt.err, "Test")
			got := outputBuffer.String()
			assert.Contains(got, tt.want, "handleErr() = %v, want = %v", got, tt.want)
		})
	}
}

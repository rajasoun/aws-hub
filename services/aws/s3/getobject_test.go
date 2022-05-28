package s3

import (
	"bytes"
	"context"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

var bucketContent = "this is the body foo bar baz"

type MockGetObjectAPI func(ctx context.Context,
	params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (mock MockGetObjectAPI) GetObject(ctx context.Context,
	params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return mock(ctx, params, optFns...)
}

func mockAPIClient() S3GetObjectAPI {
	return MockGetObjectAPI(
		func(ctx context.Context,
			params *s3.GetObjectInput,
			optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
			result := &s3.GetObjectOutput{
				Body: ioutil.NopCloser(bytes.NewReader([]byte(bucketContent))),
			}
			return result, nil
		})
}

func TestGetObjectFromS3(t *testing.T) {
	cases := []struct {
		client   func() S3GetObjectAPI
		bucket   string
		key      string
		expected []byte
		want     int
	}{
		{
			client:   func() S3GetObjectAPI { return mockAPIClient() },
			bucket:   "testBucket",
			key:      "testKey",
			expected: []byte(bucketContent),
			want:     0,
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert := assert.New(t)
			ctx := context.TODO()
			content, err := GetObjectFromS3(ctx, tt.client(), tt.bucket, tt.key)
			assert.NoError(err, "expect no error, got %v", err)
			got := bytes.Compare(tt.expected, content)
			assert.Equal(tt.want, got, "got byte recived = %v, want = %v", got, tt.want)
		})
	}
}

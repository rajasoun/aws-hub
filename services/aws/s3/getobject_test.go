package s3

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

var bucketContent = "this is the body foo bar baz"

type mockGetObjectAPI func(ctx context.Context,
	params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (mock mockGetObjectAPI) GetObject(ctx context.Context,
	params *s3.GetObjectInput,
	optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return mock(ctx, params, optFns...)
}

func mockS3GetObjectAPI(t *testing.T) S3GetObjectAPI {
	return mockGetObjectAPI(func(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
		result := &s3.GetObjectOutput{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(bucketContent))),
		}
		return result, nil
	})
}

func TestGetObjectFromS3(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	t.Parallel()

	cases := []struct {
		name   string
		client func(t *testing.T) S3GetObjectAPI
		bucket string
		key    string
		expect []byte
	}{
		{
			name: "Check Get Object From S3",
			client: func(t *testing.T) S3GetObjectAPI {
				return mockS3GetObjectAPI(t)
			},
			bucket: "testBucket",
			key:    "testKey",
			expect: []byte(bucketContent),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			content, err := GetObjectFromS3(ctx, tt.client(t), tt.bucket, tt.key)
			assert.NoError(err, "expect no error, got %v", err)
			assert.Equal(tt.expect, content, "expect %v, got %v", tt.expect, content)
		})
	}
}

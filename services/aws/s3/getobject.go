package s3

import (
	"context"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Interface for Amazon S3 API operations required by GetObjectFromS3 function
type S3GetObjectAPI interface {
	GetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

//Amazon S3 client’s GetObject method,
func GetObjectFromS3(ctx context.Context,
	client S3GetObjectAPI, bucket, key string) ([]byte, error) {
	object, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()
	return ioutil.ReadAll(object.Body)
}

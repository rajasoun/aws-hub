package s3

import (
	"context"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3GetObjectAPI interface {
	GetObject(ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

func GetObjectFromS3(ctx context.Context,
	api S3GetObjectAPI, bucket, key string) ([]byte, error) {
	object, err := api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()
	return ioutil.ReadAll(object.Body)
}

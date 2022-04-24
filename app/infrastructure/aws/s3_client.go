package aws

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error)
	ListObjects(ctx context.Context, bucket string, key string) ([]S3Obj, error)
	PutObject(ctx context.Context, bucket string, key string, body io.Reader) error
}

type s3Client struct {
	cli *s3.Client
}

type S3Obj struct {
	Bucket string
	Key    string
}

func NewS3Client(config aws.Config) *s3Client {
	return &s3Client{
		cli: s3.NewFromConfig(config),
	}
}

func (c s3Client) GetObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error) {
	res, err := c.cli.GetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (c s3Client) ListObjects(ctx context.Context, bucket string, key string) ([]S3Obj, error) {
	objects := make([]S3Obj, 0)
	var nextToken *string
	for {
		res, err := c.cli.ListObjectsV2(
			ctx,
			&s3.ListObjectsV2Input{
				Bucket:            aws.String(bucket),
				Prefix:            aws.String(key),
				ContinuationToken: nextToken,
			},
		)
		if err != nil {
			return nil, err
		}
		for _, content := range res.Contents {
			if *content.Key == key {
				continue
			}
			objects = append(objects, S3Obj{Bucket: bucket, Key: *content.Key})
		}
		if res.NextContinuationToken == nil {
			break
		}
		nextToken = res.NextContinuationToken
	}
	return objects, nil
}

func (c s3Client) PutObject(
	ctx context.Context,
	bucket string,
	key string,
	body io.Reader,
) error {
	output, err := c.cli.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   body,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}

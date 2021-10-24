package services

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	dep "github.com/blitzshare/blitzshare.fileshare.api/app/dependencies"
	"github.com/google/uuid"
)

type ShrareUrlConfig struct {
	Url          string
	ExpirationMs int64
}

func constructShareS3Key(deps *dep.Dependencies) string {
	return fmt.Sprintf("%s/%s", deps.Config.Settings.S3BucketUploadKey, uuid.New())
}

func GetPresignedUrl(deps *dep.Dependencies) ShrareUrlConfig {
	settings := deps.Config.Settings
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(settings.S3BucketRegion)},
	)
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(settings.S3BucketName),
		Key:    aws.String(constructShareS3Key(deps)),
	})
	expiration := 15 * time.Minute
	url, err := req.Presign(expiration)
	if err != nil {
		log.Println("Failed to sign request", err)
	}
	expirationMs := expiration.Milliseconds()
	return ShrareUrlConfig{
		Url:          url,
		ExpirationMs: expirationMs,
	}
}

package aws

import (
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dop251/goja"
	"time"
)

type s3Object struct {
	session    *session.Session
	bucketName string
}

func (s *s3Object) PreSignDownload(fileName string) string {
	uploader := s3manager.NewDownloader(s.session)

	var expireDuration = time.Hour * 24 * 7

	s3Req, _ := uploader.S3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: util.Pointer(s.bucketName),
		Key:    util.Pointer(fileName),
	})

	url, _, err := s3Req.PresignRequest(expireDuration)

	if err != nil {
		panic(err)
	}

	return url
}

func (s *s3Object) PreSignUpload(fileName string) string {
	uploader := s3manager.NewUploader(s.session)

	var expireDuration = time.Hour * 24 * 7

	s3Req, _ := uploader.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket: util.Pointer(s.bucketName),
		Key:    util.Pointer(fileName),
	})

	url, _, err := s3Req.PresignRequest(expireDuration)

	if err != nil {
		panic(err)
	}

	return url
}

func (s *s3Object) Delete(fileName string) {
	uploader := s3manager.NewUploader(s.session)

	_, err := uploader.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: util.Pointer(s.bucketName),
		Key:    util.Pointer(fileName),
	})

	if err != nil {
		panic(err)
	}
}

type Config struct {
	Region      string
	Endpoint    string
	Credentials struct {
		AccessKeyID     string
		SecretAccessKey string
	}
	S3ForcePathStyle bool
}

func prepareAwsConfig(config Config) *aws.Config {
	awsConfig := &aws.Config{}
	awsConfig.Region = util.Pointer(config.Region)
	awsConfig.Endpoint = util.Pointer(config.Endpoint)
	awsConfig.Credentials = credentials.NewStaticCredentials(config.Credentials.AccessKeyID, config.Credentials.SecretAccessKey, "")
	awsConfig.S3ForcePathStyle = util.Pointer(config.S3ForcePathStyle)

	return awsConfig
}

func Register(vm *goja.Runtime) error {
	return vm.Set("s3", func(config Config, bucketName string) map[string]interface{} {
		sess := session.Must(session.NewSession(prepareAwsConfig(config)))
		obj := &s3Object{session: sess, bucketName: bucketName}

		return map[string]interface{}{
			"preSignDownload": obj.PreSignDownload,
			"preSignUpload":   obj.PreSignUpload,
			"delete":          obj.Delete,
		}
	})
}

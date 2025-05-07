package pkg

import (
	"context"
	model2 "github.com/apibrew/apibrew/modules/storage/pkg/model"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

type fileProcessor struct {
	api            api.Interface
	fileRepository api.Repository[*model2.File]
	client         *minio.Client
	bucketName     string
}

func (t fileProcessor) Mapper() Mapper[*model2.File] {
	return model2.FileMapperInstance
}

func (t fileProcessor) Register(entity *model2.File) error {
	return t.presignUrls(entity)
}

func (t fileProcessor) Update(entity *model2.File) error {
	return t.presignUrls(entity)
}

func (t fileProcessor) UnRegister(entity *model2.File) error {
	return t.client.RemoveObject(context.Background(), t.bucketName, entity.Name, minio.RemoveObjectOptions{})
}

func (t fileProcessor) presignUrls(entity *model2.File) error {
	// Set expiration time for the presigned URL
	expiryTime := time.Duration(24*7) * time.Hour // URL valid for 24 hours

	// Presign a download URL
	presignedURL, err := t.client.PresignedGetObject(context.Background(), t.bucketName, entity.Name, expiryTime, nil)
	if err != nil {
		return err
	}
	log.Println("Presigned download URL:", presignedURL)

	expiryTime = time.Duration(24*7) * time.Hour // URL valid for 24 hours

	// Presign an upload URL
	presignedUploadURL, err := t.client.PresignedPutObject(context.Background(), t.bucketName, entity.Name, expiryTime)
	if err != nil {
		return err
	}
	log.Println("Presigned upload URL:", presignedUploadURL)

	entity.DownloadUrl = presignedURL.String()
	entity.UploadUrl = presignedUploadURL.String()

	return nil
}

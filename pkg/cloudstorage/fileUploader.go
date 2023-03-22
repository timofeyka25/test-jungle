package cloudstorage

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
)

type FileUploader struct {
	bucketName    string
	storageClient *storage.Client
}

func NewFileUploader(cfg Config) (*FileUploader, error) {
	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile(cfg.CloudCredentialsPath))
	if err != nil {
		return nil, err
	}

	return newFileUploader(cfg, storageClient), nil
}

func newFileUploader(
	cfg Config,
	storageClient *storage.Client,
) *FileUploader {
	return &FileUploader{
		bucketName:    cfg.ImagesBucketName,
		storageClient: storageClient,
	}
}

func (u *FileUploader) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	sw := u.storageClient.Bucket(u.bucketName).Object(file.Filename).NewWriter(ctx)

	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}

	if _, err = io.Copy(sw, fileContent); err != nil {
		return "", err
	}

	if err = sw.Close(); err != nil {
		return "", err
	}

	return sw.Attrs().MediaLink, nil
}

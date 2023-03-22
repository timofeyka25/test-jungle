package drivestorage

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"mime/multipart"
)

type FileUploader struct {
	driveService *drive.Service
}

func NewFileUploader(cfg Config) (*FileUploader, error) {
	driveService, err := drive.NewService(context.Background(), option.WithCredentialsFile(cfg.CloudCredentialsPath))
	if err != nil {
		return nil, err
	}
	return &FileUploader{driveService: driveService}, nil
}

func (u *FileUploader) Upload(file *multipart.FileHeader) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}(fileContent)

	driveFile := &drive.File{
		Name: file.Filename,
	}
	uploadedFile, err := u.driveService.Files.Create(driveFile).Media(fileContent).Do()
	if err != nil {
		return "", err
	}

	_, err = u.driveService.Permissions.Create(uploadedFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do()

	link := fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s", uploadedFile.Id)
	return link, nil
}

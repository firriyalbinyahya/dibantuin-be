package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadService struct {
	UploadDir string
}

func NewUploadService(uploadDir string) *UploadService {
	os.MkdirAll(uploadDir, os.ModePerm)
	return &UploadService{UploadDir: uploadDir}
}

func (us *UploadService) SavePhoto(file *multipart.FileHeader) (string, error) {
	// Validasi ukuran file (foto max 2MB)
	const maxPhotoSize = 2 << 20 // 2MB
	if file.Size > maxPhotoSize {
		return "", errors.New("file size exceeds 2MB")
	}

	// Validasi ekstensi file
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		return "", errors.New("invalid file type")
	}
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	filePath := filepath.Join(us.UploadDir, filename)
	return filePath, nil
}

func (us *UploadService) SaveDocument(file *multipart.FileHeader) (string, error) {
	const maxDocSize = 5 << 20 // 5MB
	if file.Size > maxDocSize {
		return "", errors.New("file size exceeds 5MB")
	}

	allowedDocs := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}
	ext := filepath.Ext(file.Filename)
	if !allowedDocs[ext] {
		return "", errors.New("invalid file type")
	}
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	filePath := filepath.Join(us.UploadDir, filename)
	return filePath, nil
}

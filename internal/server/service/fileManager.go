package service

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CreateStorageUser(dirPath string, id int64) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func UploadFile(dirPath string, id int64, name string, data []byte) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	// Write data to file
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(dirPath string, id int64, name string) ([]byte, error) {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	reader := strings.NewReader(path)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RemoveFile(dirPath string, id int64, name string) error {
	userId := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userId, "/", name)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

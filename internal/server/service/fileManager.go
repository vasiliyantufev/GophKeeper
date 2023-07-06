package service

import (
	"os"
	"path/filepath"
	"strconv"
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

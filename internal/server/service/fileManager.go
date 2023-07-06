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

//func UploadStorageUser(dirPath string) error {
//	err := os.MkdirAll(dirPath, 0755)
//	if err != nil {
//		return err
//	}
//	return nil
//}

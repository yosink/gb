package file

import (
	"fmt"
	"os"
)

func CheckNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func CheckPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

func IfNotExistMkdir(path string) error {
	if CheckNotExist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func MustOpen(filename, filepath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("getRootDir error : %v", err)
	}

	src := dir + "/" + filepath

	if CheckPermission(src) {
		return nil, fmt.Errorf("permission is denied.error path:%s", src)
	}
	if err = IfNotExistMkdir(src); err != nil {
		return nil, err
	}
	return os.OpenFile(src+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
}

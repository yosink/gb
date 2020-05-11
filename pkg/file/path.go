package file

import (
	"os"
	"path/filepath"
)

func GetWebDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func GetRootDir() (string, error) {
	webDir, err := GetWebDir()
	if err != nil {
		return "", err
	}
	return filepath.Dir(webDir), nil
}
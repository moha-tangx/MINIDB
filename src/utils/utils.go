package utils

import (
	"os"
	"strings"
)

func GetErrMsg(err error) string {
	errArr := strings.Split(err.Error(), ":")
	return errArr[len(errArr)-1]
}

func GetFiles(path string) ([]os.DirEntry, error) {
	var Files []os.DirEntry
	entries, ReadErr := os.ReadDir(path)
	if ReadErr != nil {
		print(ReadErr.Error())
		return nil, ReadErr
	}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			Files = append(Files, entry)
		}
	}
	return Files, nil
}

func GetDirs(path string) ([]os.DirEntry, error) {
	var Dirs []os.DirEntry
	entries, ReadErr := os.ReadDir(path)
	if ReadErr != nil {
		print(ReadErr.Error())
		return nil, ReadErr
	}
	for _, entry := range entries {
		if entry.Type().IsDir() {
			Dirs = append(Dirs, entry)
		}
	}
	return Dirs, nil
}

package utils

import (
	"MINIDB/src/objects"
	"encoding/json"
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

func GetDBConfig() (*objects.ConfigFile, error) {
	//note: change the config file location before production
	buff, err := os.ReadFile("C:/Users/Muham/Code/GO/src/MINIDB/test/minidb.config.json")
	if err != nil {
		println("could not read config file")
		return nil, err
	}
	var configFile = new(objects.ConfigFile)
	json.Unmarshal(buff, configFile)
	return configFile, nil
}

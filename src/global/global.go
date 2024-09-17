package global

import (
	"MINIDB/src/objects"
	"encoding/json"
	"os"
	"path"
	"runtime"
)

var DataPath = getDataPath()
var DBInUSe *objects.DATABASE = nil

var env = map[string]string{"windows": "C://", "linux": "/usr/local"}
var minidb_path, exits = os.LookupEnv("minidbpath")
var OS = runtime.GOOS

func GetDBConfig() (*objects.ConfigFile, error) {
	if !exits {
		minidb_path = path.Join(env[OS], "minidb")
	}
	config_file_path := path.Join(minidb_path, "bin/minidb.config.json")
	buff, err := os.ReadFile(config_file_path)
	if err != nil {
		println("could not read config file")
		return nil, err
	}
	var configFile = new(objects.ConfigFile)
	json.Unmarshal(buff, configFile)
	return configFile, nil
}

func getDataPath() (DataPath string) {
	DataPath = path.Join(env[OS], "minidb/data")
	configFile, err := GetDBConfig()
	if err != nil {
		return
	}
	DataPath = configFile.Storage.DataPath
	return
}

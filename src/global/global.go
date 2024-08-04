package global

import (
	"MINIDB/src/objects"
	"MINIDB/src/utils"
	"os"
	"path"
)

// var QueryReturn = new(objects.ActionReturn)

var DataPath = getDataPath()
var DBInUSe *objects.DATABASE = nil

func getDataPath() (DataPath string) {
	configFile, err := utils.GetDBConfig()
	if err != nil {
		rootDir, hasRootDir := os.LookupEnv("HOMEDRIVE")
		if hasRootDir {
			DataPath = path.Join(rootDir, "minidb/data")
			return
		}
	}
	DataPath = configFile.Storage.DataPath
	return
}

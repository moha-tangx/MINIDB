package main

import (
	"MINIDB/src/objects"
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

var rootDir, hasHome = os.LookupEnv("HOMEDRIVE")

func main() {
	fmt.Println(
		`
	you are about to install mini db press Enter to proceed or X to cancel
	`,
	)

	reader := bufio.NewReader(os.Stdin)

	proceedInstallation, err := reader.ReadString('\n')
	if err != nil {
		print(err.Error())
		return
	}

	proceedInstallation = strings.TrimSpace(proceedInstallation)
	if proceedInstallation == "x" || proceedInstallation == "X" {
		os.Exit(0)
		return
	}

	fmt.Println(
		`
	----------------------------------------------------
	specify installation path or press enter for default
		`,
		"\033[031m white spaces not allowed \033[0m",
		`
	-----------------------------------------------------
	`,
	)
	installationPath, err := reader.ReadString('\n')
	if err != nil {
		print(err.Error())
		return
	}

	// if user did not provide an installation path we set it to users root directory
	installationPath = strings.TrimSpace(installationPath)
	if installationPath == "" {
		if !hasHome {
			println("$HOMEDRIVE env variable not found")
			return
		}
		installationPath = rootDir
	}

	installationPath = path.Clean(installationPath)
	// path to main directory of minidb
	mainDirPath := path.Join(installationPath, "minidb")

	// path to bin directory where programme is installed
	binPath := path.Join(mainDirPath, "bin")

	if err := os.MkdirAll(binPath, os.ModeDir); err != nil {
		println("unexpected error, could not create dir /bin")
		println(err.Error())
		return
	}

	// path to log directory where log file are stored
	logPath := path.Join(mainDirPath, "log")
	if err := os.MkdirAll(logPath, os.ModeDir); err != nil {
		println("unexpected error, could not create dir /log")
		println(err.Error())
		return
	}

	fmt.Println(
		`
	-------------------------------------------------
	specify to store data or press enter for default
		`,
		"\033[031m white spaces not allowed \033[0m",
		`
	--------------------------------------------------
	`,
	)
	dataPath, err := reader.ReadString('\n')
	if err != nil {
		print(err.Error())
		return
	}
	dataPath = strings.TrimSpace(dataPath)
	if dataPath == "" {
		dataPath = path.Join(mainDirPath, "data")
	}

	dataPath = path.Clean(dataPath)
	// create directory to store data
	if err := os.MkdirAll(dataPath, os.ModeDir); err != nil {
		println("unexpected error, could not create dir /data")
		println(err.Error())
		return
	}
	// path to database configuration file
	config_FIle_Path := path.Join(binPath, "minidb.config.json")

	// create database configuration file
	Config_file, configErr := os.OpenFile(config_FIle_Path, os.O_CREATE, os.ModePerm)
	if configErr != nil {
		println("could not create configuration fle")
		return
	}
	defer Config_file.Close()

	var dbConfig objects.ConfigFile
	dbConfig.Storage.DataPath = dataPath
	dbConfig.Logs.Path = logPath
	dbConfig.Net.Ip = "127.0.0.1"
	dbConfig.Net.Port = "17390"

	buff, MarshErr := json.MarshalIndent(dbConfig, "", "  ")

	if MarshErr != nil {
		println("could not parse db config")
	}

	if _, err := Config_file.Write(buff); err != nil {
		println("could not write bytes to file")
		return
	}

	// path to database configuration file
	log_FIle_Path := path.Join(logPath, "mini.log")

	// create database configuration file
	Log_file, logErr := os.Create(log_FIle_Path)
	Log_file.Chmod(fs.FileMode(os.O_RDONLY))
	if logErr != nil {
		println("could not create configuration fle")
		return
	}
	defer Log_file.Close()

	fmt.Println(
		"\033[32m",
		`
	INSTALLING ....
	`, "\033[0m")
	fmt.Println(
		`
	done
	`)
}

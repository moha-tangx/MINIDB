package queries

import (
	"MINIDB/src/global"
	"MINIDB/src/objects"
	"MINIDB/src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var DBInUSe = global.DBInUSe
var DataPath = global.DataPath

func CreateDB(DBName string) (string, error) {
	if err := os.Mkdir(DataPath+"/"+DBName, os.ModePerm); err != nil {
		fmt.Printf("could not create database %s\n", DBName)
		errMsg := utils.GetErrMsg(err)
		if strings.HasSuffix(errMsg, "already exists.") {
			println("database already exists")
			return "", &objects.UserError{
				Type:    "duplicate error",
				Root:    nil,
				Message: "database already exists",
			}
		}
		println("system error")
		return "", &objects.UserError{
			Type:    "systemError",
			Root:    err,
			Message: "could not create database",
		}
	}
	println("database created")
	return "database created", nil
}

func DropDB(DBName string) (string, error) {
	DBPath := path.Join(DataPath, DBName)
	entries, _ := utils.GetDirs(DataPath)
	for _, entry := range entries {
		if entry.Name() == DBName {
			// if match is found we go ahead and remove directory
			goto remove
		}
	}
	println("could not drop database")
	println("database not found")
	return "", &objects.UserError{
		Type:    "objectNotFound",
		Root:    nil,
		Message: "could not drop database \n database not found",
	}

remove:
	if err := os.RemoveAll(DBPath); err != nil {
		fmt.Printf("could not drop database %s \n", DBName)
		println(err.Error())
		return "", &objects.UserError{
			Type:    "system error",
			Root:    err,
			Message: "could not drop database",
		}
	}
	println("database dropped")
	return "database dropped", nil
}

func GetDBs() ([]os.DirEntry, error) {
	dbs := []os.DirEntry{}
	entries, readErr := os.ReadDir(DataPath)
	if readErr != nil {
		fmt.Println("could not retrieve databases")
		return nil, &objects.UserError{
			Type:    "systemError",
			Root:    readErr,
			Message: "could not retrieve databases",
		}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dbs = append(dbs, entry)
		}
	}
	return dbs, nil
}

func GetDBInUse() (*objects.DATABASE, error) {
	if DBInUSe != nil {
		return DBInUSe, nil
	}
	databases, GetErr := GetDBs()
	if GetErr != nil {
		println("could not retrieve databases")
		print(GetErr.Error())
		return nil, GetErr
	}
	if len(databases) < 1 {
		return nil, GetErr
	}
	firstDB := databases[0].Name()
	SetDBInUse(firstDB)
	return DBInUSe, nil
}

func ShowDBs() ([]string, error) {
	dbs, err := GetDBs()
	dbNames := []string{}
	if err != nil {
		println("could not get databases")
		println(err)
		return nil, err
	}
	for _, db := range dbs {
		dbNames = append(dbNames, db.Name())
		println(db.Name())
	}
	// println()
	return dbNames, nil
}

// sets the database to be the one in use
func SetDBInUse(DBName string) {
	DBInUSe = new(objects.DATABASE)
	DBInUSe.Path = path.Join(DataPath, DBName)
	DBInUSe.Name = DBName
	var collections []*objects.COLLECTION
	files, GetErr := utils.GetFiles(DBInUSe.Path)
	if GetErr != nil {
		return
	}
	for _, file := range files {
		collection := objects.COLLECTION{
			Name:     strings.Split(file.Name(), ".")[0],
			Path:     path.Join(DBInUSe.Path, file.Name()),
			Database: DBInUSe,
		}
		collections = append(collections, &collection)
	}
	DBInUSe.Collections = collections
}

func UseDB(DBName string) (string, error) {
	dbs, GetErr := GetDBs()
	if GetErr != nil {
		return "", &objects.UserError{Type: "systemError", Root: GetErr, Message: "could not switch database"}
	}
	for _, db := range dbs {
		if db.Name() == DBName {
			SetDBInUse(DBName)
			fmt.Printf("switched to %v \n", db.Name())
			return "switched to " + db.Name() + "\n", nil
		}
	}
	println("database with the specified name could not be found")
	return "", &objects.UserError{Type: "objectNotFound", Root: nil, Message: "database with the specified name could not be found"}
}
func CreateCollection(CollectionName string) (string, error) {
	filename := CollectionName + ".json"
	DBinUSe, err := GetDBInUse()
	if err != nil {
		return "could not identify database in use", err
	}
	collectionPath := path.Join(DBinUSe.Path, filename)
	collections, GetErr := GetCollections()
	if GetErr != nil {
		return "could not get collection", &objects.UserError{Type: "systemError", Root: GetErr, Message: "could not create const name = value"}
	}
	for _, collection := range collections {
		if collection.Name() == filename {
			println("could not create collection")
			println("collection already exists")
			return "", &objects.UserError{Type: "duplicateError", Root: nil, Message: "collection already exists"}
		}
	}
	collection, CreateErr := os.Create(collectionPath)
	if CreateErr != nil {
		println("could not create collection")
		return "could not create collection", CreateErr
	}
	defer collection.Close()
	println("collection created successfully")
	return "collection created successfully", nil
}

func DropCollection(collectionName string) error {
	fileName := collectionName + ".json"
	DBinUSe, GetErr := GetDBInUse()
	if GetErr != nil {
		return GetErr
	}
	filePath := path.Join(DBinUSe.Path, fileName)

	if err := os.Remove(filePath); err != nil {
		println("could not drop collection")
		if strings.HasSuffix(err.Error(), "cannot find the file specified.") {
			println("collection not found")
			return errors.New("collection not found")
		}
		println("system error")
		return err
	}
	println("collection dropped")
	return nil
}

func GetCollections() ([]os.DirEntry, error) {
	var collections []os.DirEntry
	DBInUSe, getDBerr := GetDBInUse()
	if getDBerr != nil {
		return nil, getDBerr
	}
	if DBInUSe == nil {
		return nil, nil
	}
	entries, err := utils.GetFiles(DBInUSe.Path)
	if err != nil {
		println("could not get collections")
		// print(err.Error())
		return nil, err
	}
	for _, entry := range entries {
		filepath := path.Join(DBInUSe.Path, entry.Name())
		if path.Ext(filepath) == ".json" {
			collections = append(collections, entry)
		}
	}
	return collections, nil
}

func ShowCollections() ([]string, error) {
	collections, getErr := GetCollections()
	collectionNames := []string{}
	if getErr != nil {
		return nil, getErr
	}
	for _, collection := range collections {
		name := strings.Split(collection.Name(), ".")[0]
		collectionNames = append(collectionNames, name)
		println(name)
	}
	return collectionNames, nil
}

func ClearConsole() {
	clearCommand := "clear"
	if runtime.GOOS == "windows" {
		clearCommand = "cls"
	}
	c := exec.Command(clearCommand)
	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}

func InsertDocument(collectionName string, document string) error {
	collectionName = collectionName + ".json"
	if !json.Valid([]byte(document)) {
		println("could not 1")
		return errors.New("invalid json")
	}
	collections, err := GetCollections()
	if err != nil {
		println("could not 2")
		return err
	}
	for _, collection := range collections {
		if collection.Name() == collectionName {
			collectionFilePath := path.Join(DBInUSe.Path, collectionName)
			file, err := os.OpenFile(collectionFilePath, os.O_APPEND, 0)
			if err != nil {
				println("could not 3")
				return err
			}
			if _, err := file.WriteString(document + ","); err != nil {
				println("could not 4")
				return err
			}
			return nil
		}
	}
	return errors.New("collection not found")
}

func Exit() {
	os.Exit(0)
}

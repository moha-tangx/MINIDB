package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)


type DATABASE struct{
	DBName string
	DBPath string
	collections []string
}

// check for users os before navigation
const dataDIr = "C://users/muham/code/go/src/minidb/test"

var DBinUSePath = ""

func main() {

}

func getErrMsg(err error) string{
	errArr := strings.Split(err.Error(), ":")
	return errArr[len(errArr)-1]
}

func getFiles(path string)([]os.DirEntry,error){
	var Files[]os.DirEntry
	entries, ReadErr := os.ReadDir(path);
	if ReadErr != nil{
		print(ReadErr.Error())
		return nil, ReadErr
	}
	for _,entry := range entries{
		if entry.Type().IsRegular(){
			Files = append(Files, entry)
		}
	}
	return Files,nil
}

func getDirs(path string)([]os.DirEntry,error){
	var Dirs[]os.DirEntry
	entries, ReadErr := os.ReadDir(path);
	if ReadErr != nil{
		print(ReadErr.Error())
		return nil, ReadErr
	}
	for _,entry := range entries{
		if entry.Type().IsDir(){
			Dirs = append(Dirs, entry)
		}
	}
	return Dirs,nil
}

func GEtDBInUse()(*DATABASE,error){
	if(DBinUSePath == ""){
		databases, GetErr:= GetDBs()
		if GetErr != nil{
			println("could not retrieve databases")
			print(GetErr.Error())
			return nil,GetErr
		}
		firstDBName := databases[0].Name()
		DBinUSePath = path.Join(dataDIr,firstDBName)
	}
	var collections []string

	files, GetErr := getFiles(DBinUSePath)
	if GetErr != nil{
		return nil,GetErr
	}
	for _,file  := range files {
		collections = append(collections, file.Name())
	}
	dbInUse := DATABASE{
		DBName: path.Dir(DBinUSePath),
		DBPath: DBinUSePath,
		collections: collections,
	}
	return &dbInUse,nil
}

func CreateDB(DBName string)error{
	if err :=	os.Mkdir(dataDIr+"/"+DBName,os.ModePerm); err != nil{
		fmt.Printf("could not create database %s\n", DBName)
		errMsg := getErrMsg(err)
		if strings.HasSuffix(errMsg,"already exists.",) {
			println("database already exists")
			return errors.New("database already exists")
		}
		println("system error")
		return errors.New(err.Error())
	}
	println("database created")
	return nil
}

func DropDB(DBName string)error{
	DBPath := path.Join(dataDIr,DBName)
	entries, _ := getDirs(dataDIr)
	for _, entry := range entries {
		if entry.Name() == DBName{
			// if match is found we go ahead and remove directory
			goto remove
		}
	}
	println("could not drop database")
	println("database not found")
	return errors.New("database not found")

	remove: if err := os.RemoveAll(DBPath);err !=nil{
		fmt.Printf("could not drop database %s \n", DBName)
		println(err.Error())
		return err
	}
	println("database dropped")
	return nil
}

func GetDBs()([]os.DirEntry,  error){
	dbs := []fs.DirEntry{}
	entries, readErr := os.ReadDir(dataDIr)
	if readErr != nil{
		fmt.Println("could not get databases")
		return nil ,readErr
	}
	for _,entry  := range entries {
		if entry.IsDir(){
			dbs = append(dbs, entry)
		}
	}
	return dbs, nil;
}

func ShowDBs() error{
	dbs, err := GetDBs()
	if err != nil {
		println("could not get databases")
		println(err)
		return err
	}
	for _,db  := range dbs {
		println(db.Name())
	}
	println()
	return nil
}

func UseDB(DBName string)error{
	dbs, GetErr := GetDBs();
	if GetErr != nil{
		return GetErr
	}
	for _,db  := range dbs {
		if db.Name() == DBName{
			DBinUSePath = path.Join(dataDIr,DBName);
			fmt.Printf("switched to %v \n", db.Name())
			return nil
		}
	}
	println("database with the specified name could not be found")
	return errors.New("database with the specified name could not be found ")
}


func CreateCollection(CollectionName string) error{
	filename := CollectionName+".json"
	DBinUSe,_ := GEtDBInUse()
	collectionPath := path.Join(DBinUSe.DBPath,filename)
	collections, GetErr := GetCollections()
	if(GetErr != nil){
		return GetErr
	}
	for _, collection := range collections {
		if collection.Name() == filename {
			println("could not create collection")
			println("collection already exists")
			return errors.New("collection already exists")
		}
	}
	 collection , CreateErr := os.Create(collectionPath)
	if CreateErr != nil{
		println("could not create collection")
		return CreateErr
	}
	defer collection.Close()
	println("collection created successfully")
	return nil
}

func DropCollection(collectionName string)error{
	fileName := collectionName+".json"
	DBinUSe,_ := GEtDBInUse()
	filePath := path.Join(DBinUSe.DBPath,fileName)

	if err :=  os.Remove(filePath); err != nil{
		println("could not drop collection")
		if (strings.HasSuffix(err.Error(),"cannot find the file specified.")){
			println("collection not found")
			return errors.New("collection not found")
		}
		println("system error")
		return err
	}
	println("collection dropped")
	return nil
}

func GetCollections()( []os.DirEntry, error){
   var collections[]os.DirEntry
    DBInUSe,_ := GEtDBInUse()
	entries,err := getFiles(DBInUSe.DBPath);
	if err != nil{
		print(err.Error())
		return nil, err
	}
	for _, entry := range entries{
		filepath := path.Join(DBinUSePath,entry.Name()) 
		if path.Ext(filepath) == ".json"{
			collections = append(collections, entry)
		}
	}
	return collections,nil
}

func ShowCollections()error{
	collections, getErr := GetCollections();
	if getErr != nil{
		return getErr
	}
	for _,collection  := range collections {
		name :=	strings.Split(collection.Name(), ".")[0]
		println(name)
	}
	println()
	return nil
}

func Exit(){
	os.Exit(0)
}
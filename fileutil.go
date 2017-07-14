package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"errors"
)

type FileUtil struct {}

func (fileutil FileUtil) Tree(dir string) ([]string, error) {
	if dir == ""{
		errMessage := "Directory param must not be an empty string"
		log.Fatal(errMessage)
		return nil, errors.New(errMessage)
	}
	log.Println("Building directory tree for:", dir)

	var list []string

	err := fileutil.mountTreeOfFilesAndDirs(dir, &list)

	// Removing the root path from files and dirs name
	for index, filename := range list{
		list[index] = strings.Replace(filename, dir, "", -1)
	}

	return list, err

}

func (fileutil FileUtil) Open(fileName string) (string, error){
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error while opening file,\n", err)
		return "", err
	}
	return string(file), nil
}

func (fileutil FileUtil) mountTreeOfFilesAndDirs(dir string, list *[]string) error{

	currentList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Error while listing directories.\n", err)
		return err
	}

	for _, file := range currentList {

		fileKind := "${F}"
		if file.IsDir(){
			fileKind = "${D}"
		}

		completeFileName := fmt.Sprintf("%s/%s", dir, file.Name())
		outputFileName := fmt.Sprintf("%s%s/%s", fileKind, dir, file.Name())
		*list = append(*list, outputFileName)

		if file.IsDir() {
			fileutil.mountTreeOfFilesAndDirs(completeFileName, list)
		}
	}
	return nil

}
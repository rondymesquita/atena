package main

import (
	"fmt"
	"log"
	"strings"
	"regexp"
)

func NewAuditor(config Config) *Auditor{
	return &Auditor{config}
}

type Auditor struct{
	config Config
}

func (auditor Auditor) Start() {
	auditor.config.Load()
	//fmt.Println(auditor.config)

	f := FileUtil{}
	list, _ := f.Tree(auditor.config.Management.Root)
	fmt.Println("final list here: ", len(list))
	//log.Println(list[82])
	for index, el := range list{
		log.Println(index, el)
	}
	for index, el := range auditor.config.Management.Directories{
		log.Println(index, el)
	}

	listResult , err := auditor.validateManagementRepository(list)
	if err != nil {
		log.Fatal("err on validate, ", err)
	}
	for _, result := range listResult{
		log.Println(result)
	}
}

func(auditor Auditor) validateManagementRepository(list []string) ([]string, error){
	for _, completeFileName := range list{
		for _, directory := range auditor.config.Management.Directories{

			if strings.Contains(completeFileName, "${F}") {
				handleFile(completeFileName, directory)
			} else if strings.Contains(completeFileName, "${D}") {
				handleDirectory(completeFileName, directory)
			}

		}
	}
	return nil, nil
}

func handleFile(completeFileName string, directory Directory){
	if strings.Contains(completeFileName, directory.Name){

		hasMatch := true
		fileName := strings.Replace(completeFileName, directory.Name, "" , -1)
		fileName = strings.Replace(fileName, "${F}/", "" , -1)

		for _, rule := range directory.Rules{

			log.Println(fileName, rule.Format)

			//match regex
			regex := fmt.Sprintf("^%s$",rule.Format)
			match, _ :=  regexp.MatchString(regex, fileName)

			if !match{
				hasMatch = match
				break
			}

		}

		log.Println(hasMatch)

	}
}

func handleDirectory(completeFileName string, directory Directory){

}


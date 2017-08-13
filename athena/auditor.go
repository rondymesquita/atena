package athena

import (
	"fmt"
	"log"
	"strings"
)

func NewAuditor(config Config, parser Parser) *Auditor {
	return &Auditor{config, parser}
}

type Auditor struct {
	config Config
	parser Parser
}

func (auditor Auditor) Start() {
	auditor.config.Load()

	f := FileUtil{}
	list, _ := f.Tree(auditor.config.Management.Root)
	fmt.Println("final list here: ", len(list))
	//log.Println(list[82])
	for index, el := range list {
		log.Println(index, el)
	}
	for index, el := range auditor.config.Management.Directories {
		log.Println(index, el)
	}

	listResult, err := auditor.validateManagementRepository(list)
	if err != nil {
		log.Fatal("err on validate, ", err)
	}
	for _, result := range listResult {
		log.Println(result)
	}
}

func (auditor Auditor) validateManagementRepository(list []string) ([]string, error) {
	for _, completeFileName := range list {
		for _, directory := range auditor.config.Management.Directories {

			if strings.Contains(completeFileName, "${F}") {
				auditor.handleFile(completeFileName, directory)
			} else if strings.Contains(completeFileName, "${D}") {
				handleDirectory(completeFileName, directory)
			}

		}
	}
	return nil, nil
}

func (auditor Auditor) handleFile(completeFileName string, directory Directory) {
	if strings.Contains(completeFileName, directory.Name) {

		hasMatch := true
		fileName := strings.Replace(completeFileName, directory.Name, "", -1)
		fileName = strings.Replace(fileName, "${F}/", "", -1)

		for _, rule := range directory.Rules {

			log.Println(fileName, rule.Pattern)

			//match regex
			//regex := fmt.Sprintf("^%s$", rule.Pattern)
			//match, _ := regexp.MatchString(regex, fileName)
			matched, err := auditor.parser.HasMatch(fileName, rule.Pattern)

			if err != nil {
				log.Fatal("err on parser, ", err)
			}

			if !matched {
				hasMatch = matched
				break
			}

		}

		log.Println(hasMatch)

	}
}

func handleDirectory(completeFileName string, directory Directory) {

}

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
	config           Config
	parser           Parser
}

func (auditor Auditor) Start() {
	auditor.config.Load()

	f := FileUtil{}
	list, _ := f.Tree(auditor.config.Management.Root)
	_, err := auditor.validateManagementRepository(list)

	if err != nil {
		log.Fatal("err on audit, ", err)
	}
}

func (auditor Auditor) validateManagementRepository(list []string) ([]string, error) {

	var machedPatterns, unmachedPatterns []string

	//for each file on the list
	for _, completeFileName := range list {

		//for each directory declared on config file
		for _, directory := range auditor.config.Management.Directories {

			if strings.Contains(completeFileName, "${F}") {
				auditor.handleFile(completeFileName, directory, &machedPatterns, &unmachedPatterns)
			} else if strings.Contains(completeFileName, "${D}") {
				handleDirectory(completeFileName, directory)
			}

		}
	}

	ui := NewUI()
	ui.PrintMatched(unmachedPatterns)
	ui.PrintUnmatched(unmachedPatterns)


	return nil, nil
}

func (auditor Auditor) handleFile(completeFileName string, directory Directory, machedPatterns *[]string, unmachedPatterns *[]string) {
	if strings.Contains(completeFileName, directory.Name) {

		//hasMatch := true
		fileName := strings.Replace(completeFileName, directory.Name, "", -1)
		fileName = strings.Replace(fileName, "${F}/", "", -1)

		//match regex
		matched, err := auditor.parser.HasMatch(fileName, directory.Pattern)

		if err != nil {
			log.Fatal("err on parser, ", err)
		}

		if !matched {
			unmachedPattern := fmt.Sprintf("\"%s\" does not match with \"%s\"", fileName, directory.Pattern)
			*unmachedPatterns = append(*unmachedPatterns, unmachedPattern)
		} else {
			machedPattern := fmt.Sprintf("=====> \"%s\" matches with \"%s\"", fileName, directory.Pattern)
			*machedPatterns = append(*machedPatterns, machedPattern)
			//auditor.machedPatterns = append(auditor.machedPatterns, machedPattern)
		}


	}
}

func handleDirectory(completeFileName string, directory Directory) {

}

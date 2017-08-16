package athena

import (
	"fmt"
	"log"
	"strings"
)

func NewAuditor(config Config, parser Parser) *Auditor {
	cli := *NewCLI()
	return &Auditor{config, parser, cli}
}

type Auditor struct {
	config Config
	parser Parser
	cli    CLI
}

func (auditor Auditor) Start() {
	auditor.config.Load()
	auditor.validate()
	//_, err := auditor.validateManagementRepository(list)

	//if err != nil {
	//	log.Fatal("err on audit, ", err)
	//}
}

func (auditor Auditor) validate() {
	config := auditor.config
	f := FileUtil{}

	var unmachedPatterns []string

	for _, directory := range auditor.config.Management.Directories {
		completeDirectoryName := fmt.Sprintf("%s%s", config.Management.Root, directory.Name)
		fileList, _ := f.List(completeDirectoryName)

		var matchList []bool
		hasMatch := true
		for _, file := range fileList {
			for _, rule := range directory.Rules {

				matched, err := auditor.parser.HasMatch(file, rule.Pattern)

				if err != nil {

				}

				matchList = append(matchList, matched)
				log.Print(completeDirectoryName, "   ", matchList)
				if !matched {
					unmachedPattern := fmt.Sprintf("\"%s\" does not match with \"%s\"", file, rule.Pattern)
					unmachedPatterns = append(unmachedPatterns, unmachedPattern)
					log.Print(unmachedPattern)
					hasMatch = false
				}

			}

		}
		log.Println(hasMatch)
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

	cli := NewCLI()

	if len(unmachedPatterns) == 0 {
		cli.Print("Directories validated with success.")
	} else {
		cli.PrintUnmatched(unmachedPatterns)
	}

	return nil, nil
}

func (auditor Auditor) handleFile(completeFileName string, directory Directory, machedPatterns *[]string, unmachedPatterns *[]string) {
	if strings.Contains(completeFileName, directory.Name) {

		//hasMatch := true
		var matchList []bool
		fileName := strings.Replace(completeFileName, directory.Name, "", -1)
		fileName = strings.Replace(fileName, "${F}/", "", -1)

		//match regex with each rule pattern
		for _, rule := range directory.Rules {

			matched, err := auditor.parser.HasMatch(fileName, rule.Pattern)
			matchList = append(matchList, matched)
			if err != nil {
				log.Fatal("err on parser, ", err)
			}

			if !matched {
				//hasMatch = matched
				unmachedPattern := fmt.Sprintf("\"%s\" does not match with \"%s\"", fileName, rule.Pattern)
				*unmachedPatterns = append(*unmachedPatterns, unmachedPattern)
			} else {
				machedPattern := fmt.Sprintf("=====> \"%s\" matches with \"%s\"", fileName, rule.Pattern)
				*machedPatterns = append(*machedPatterns, machedPattern)
			}
		}

		log.Print(matchList)

	}
}

func handleDirectory(completeFileName string, directory Directory) {

}

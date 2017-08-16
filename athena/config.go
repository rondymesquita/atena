package athena

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

func NewConfig() *Config {
	return &Config{}
}

type Config struct {
	Management Management
}

type Management struct {
	Root        string
	Directories []Directory
}

type Directory struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	Pattern string
}

func (config *Config) Test() {
	var c Config
	c.Management.Root = "Root"

	var rules []Rule
	r1 := Rule{Pattern:"xxx"}
	r2 := Rule{Pattern:"yyy"}
	rules = append(rules, r1)
	rules = append(rules, r2)

	d := Directory{"Folder1", rules}
	d2 := Directory{"Folder1", rules}
	c.Management.Directories = append(c.Management.Directories, d)
	c.Management.Directories = append(c.Management.Directories, d2)
	dump, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatal("Error while opening Configfile\n", err)
	}
	fmt.Print(string(dump))
}

func (config *Config) Load(fileName ...string) {

	if len(fileName) == 0 {
		fileName = append(fileName, "./Configfile.yml")
	}

	configfile, err := FileUtil{}.Open(fileName[0])

	if err != nil {
		log.Fatal("Error while opening Configfile\n", err)
	}
	err = yaml.Unmarshal([]byte(configfile), &config)
	if err != nil {
		log.Fatal("Error while parsing Configfile\n", err)
	}

	config.validate()
}

//TODO validate each attribute
func (config *Config) validate() {
	returnErrorWhenAttrIsNil(config.Management)
	returnErrorWhenAttrIsNil(config.Management.Root)
	//returnErrorWhenAttrIsEmpty(config.Management.Directories)
}

func returnErrorWhenAttrIsNil(attribute interface{}) error {

	if attribute == nil {
		return errors.New(fmt.Sprintf("Attribute \"%s\" was not found on Configfile.yml", attribute))
	}

	return nil
}

func returnErrorWhenAttrIsEmpty(list []interface{}) error {

	if len(list) == 0 {
		return errors.New(fmt.Sprintf("Attribute \"%s\" was not found on Configfile.yml", list))
	}

	return nil
}

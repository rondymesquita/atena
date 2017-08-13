package athena

import (
	"gopkg.in/yaml.v2"
	"log"
)

func NewConfig() *Config{
	return &Config{}
}

type Config struct{
	Management Management
}

type Management struct{
	Root string
	Directories []Directory
}

type Directory struct{
	Name string
	Rules []Rule
}

type Rule struct{
	Pattern string
}

func (config *Config) Load(fileName ...string){

	if len(fileName) == 0{
		fileName = append(fileName, "./Configfile.yml")
	}

	configfile, err := FileUtil{}.Open("./Configfile.yml")
	//log.Println(configfile)
	if err != nil{
		log.Fatal("Error while opening Configfile\n", err)
	}
	err = yaml.Unmarshal([]byte(configfile), &config)
	if err != nil{
		log.Fatal("Error while parsing Configfile\n", err)

	}
}

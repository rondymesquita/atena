package athena

import (
	"strings"
	"fmt"
	"regexp"
	"log"
)

//var Rules = map[string]string{}



func NewParser() *Parser{

	rules := map[string]string{
		//".":"\\.",
		".": "\\.",
		"<number>": "\\d*",
		"<aaaa>": "\\d{4}",
		"<mm>": "\\d{2}",
		"<dd>": "\\d{2}",
		"<*>": ".+",


	}

	return &Parser{rules}
}

type Parser struct{
	rules map[string]string
}

func (p Parser) HasMatch(value, pattern string) (bool, error){
	regexPattern := p.regexpFrom(pattern)
	log.Println(fmt.Sprintf("%s : %s : %s", value, regexPattern, pattern))
	regex := fmt.Sprintf("^%s$",regexPattern)
	matched, err :=  regexp.MatchString(regex, value)

	if !matched{
		log.Println(fmt.Sprintf("%s does not match with %s", value, pattern))
	}

	return matched, err

}

func (p Parser) regexpFrom(value string) string{
	for k, v := range p.rules{
		value = strings.Replace(value, k, v, -1)
	}

	return value
}
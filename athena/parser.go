package athena

import (
	"fmt"
	"regexp"
	"strings"
)

//var Rules = map[string]string{}



func NewParser() *Parser{

	//sanitize strings removing char that matches with regexp exp
	stringrules := map[string]string{
		".": "\\.",
	}

	regexprules := map[string]string{
		"<number>": "\\d*",
		"<yyyy>": "\\d{4}",
		"<mm>": "\\d{2}",
		"<dd>": "\\d{2}",
		"<\\*>": ".*",
	}

	return &Parser{stringrules, regexprules }
}

type Parser struct{
	stringrules map[string]string
	regexprules map[string]string
}

func (p Parser) HasMatch(value, pattern string) (bool, error){
	regexPattern := p.regexpFrom(pattern)

	//log.Println(fmt.Sprintf("%s : %s : %s", value, regexPattern, pattern))

	regex := fmt.Sprintf("^%s$",regexPattern)
	matched, err :=  regexp.MatchString(regex, value)

	return matched, err

}

func (p Parser) regexpFrom(value string) string{

	//for k, v := range p.stringrules{
	//	value = strings.Replace(value, k, v, -1)
	//}
	value = regexp.QuoteMeta(value)

	for k, v := range p.regexprules{
		value = strings.Replace(value, k, v, -1)
	}

	return value
}

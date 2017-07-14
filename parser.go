package main

import (
	"strings"
)

//var Rules = map[string]string{}



func NewParser() *Parser{

	rules := map[string]string{
		"<number>": "\\d*",
		"<aaaa>": "\\d{4}",
		"<mm>": "\\d{2}",
		"<dd>": "\\d{2}",
		"<*>": ".*",
	}

	return &Parser{rules}
}

type Parser struct{
	rules map[string]string
}

func (p Parser) RegexpFrom(value string) string{
	for k, v := range p.rules{
		value = strings.Replace(value, k, v, -1)
	}

	return value
}

func (p Parser) HasMatch(value, pattern string){

}


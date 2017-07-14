package main

var Rules = map[string]string{}



func NewParser() *Parser{

	Rules = map[string]string{
		"<number>": "\\d*",
		"<aaaa>": "\\d{4}",
		"<mm>": "\\d{2}",
		"<dd>": "\\d{2}",
		"*": ".*",
	}

	return &Parser{Rules}
}

type Parser struct{
	Rules map[string]string
}

func (p Parser) Parse(value string) string{

	return ""
}


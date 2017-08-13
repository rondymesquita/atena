package main

import "./athena"

func main(){
  config := athena.NewConfig()
  parser := athena.NewParser()
  auditor := athena.NewAuditor(*config, *parser)
  auditor.Start()
}

package main

import "atena/athena"

func main(){
  config := athena.NewConfig()
  //config.Test()
  parser := athena.NewParser()
  auditor := athena.NewAuditor(*config, *parser)
  auditor.Start()
}

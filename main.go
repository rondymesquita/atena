package main

import "atena/athena"

func main(){
  config := athena.NewConfig()
  parser := athena.NewParser()
  auditor := athena.NewAuditor(*config, *parser)
  auditor.Start()
  //fmt.Print("efsf")
}

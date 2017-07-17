package main

func main(){
  config := NewConfig()
  parser := NewParser()
  auditor := NewAuditor(*config, *parser)
  auditor.Start()
}

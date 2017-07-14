package main

func main(){
  config := NewConfig()
  auditor := NewAuditor(*config)
  auditor.Start()
}

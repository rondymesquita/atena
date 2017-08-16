package athena

import (
	"log"
	"os"
	"fmt"
	"github.com/fatih/color"
)

func NewCLI () *CLI{
	//l := log.New(os.Stdout,  fmt.Sprintf("[%s] ", ATHENA), 0)
	l := log.New(os.Stdout,  "", 0)
	return &CLI{l}
}

type CLI struct {
	log *log.Logger
}

func (cli CLI) Print(value string){
	cli.log.Println(value)
}

func (cli CLI) PrintHeader(value string){
	cli.log.Println(fmt.Sprintf("*** %s", value))
}

func (cli CLI) PrintMatched(matched []string){
	cli.PrintHeader("Matched items")

	color.Set(color.FgGreen)
	defer color.Unset()
	for _, v := range matched{
		cli.Print(v)
	}
	cli.log.Println()
}

func (cli CLI) PrintUnmatched(unmatched []string){
	cli.PrintHeader("Unmatched items")

	color.Set(color.FgRed)
	defer color.Unset()
	for _, v := range unmatched{
		cli.Print(v)
	}
	cli.log.Println()
}
package athena

import (
	"log"
	"os"
	"fmt"
	"github.com/fatih/color"
)

func NewUI () *UI{
	//l := log.New(os.Stdout,  fmt.Sprintf("[%s] ", ATHENA), 0)
	l := log.New(os.Stdout,  "", 0)
	return &UI{l}
}

type UI struct {
	log *log.Logger
}

func (ui UI) Print(value string){
	ui.log.Println(value)
}

func (ui UI) PrintHeader(value string){
	ui.log.Println(fmt.Sprintf("*** %s", value))
}

func (ui UI) PrintMatched(matched []string){
	ui.PrintHeader("Matched items")

	color.Set(color.FgGreen)
	defer color.Unset()
	for _, v := range matched{
		ui.Print(v)
	}
	ui.log.Println()
}

func (ui UI) PrintUnmatched(unmatched []string){
	ui.PrintHeader("Unmatched items")

	color.Set(color.FgRed)
	defer color.Unset()
	for _, v := range unmatched{
		ui.Print(v)
	}
	ui.log.Println()
}
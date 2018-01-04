package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
)

func main() {

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	p := NewPrinting(s)
	p.Clear()
	menu := NewMenu()
	menu.Print(p)

	p.Show()
	go menu.EventManager(p)

	<-menu.Wait

	s.Fini()
}

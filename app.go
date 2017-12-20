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

	quit := make(chan struct{})

	menu := NewMenu()
	menu.Print(p)

	p.Show()
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				case tcell.KeyUp:
					menu.Prev(p)
					p.Show()
				case tcell.KeyDown:
					menu.Next(p)
					p.Show()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	<-quit

	s.Fini()
}

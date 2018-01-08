package main

import (
	"fmt"
)

func NewTestMenu() *Menu {
	tmpcs := []Command{}
	i := 1
	for i < 5 {
		tmpc := Command{
			Title:    fmt.Sprintf("Command %d", i),
			Cli:      "foo",
			Optional: true,
		}
		tmpcs = append(tmpcs, tmpc)
		i++
	}
	m := NewMenu()
	m.Commands = tmpcs
	return m
}

func main() {
	menu := NewTestMenu()

	menu.Show()
	go menu.EventManager()
	<-menu.Wait
	menu.Quit()
}

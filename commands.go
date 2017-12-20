package main

import (
	"fmt"
)

type Command struct {
	Title string
	Cli   string
}

type Menu struct {
	Commands []Command
	Cursor   int
}

func (m *Menu) Print(p *Printing) {
	for i, c := range m.Commands {
		p.Putln(c.Title, i == m.Cursor)
	}
}

func (m *Menu) Next(p *Printing) {
	if m.Cursor < len(m.Commands)-1 {
		m.Cursor++
	} else {
		m.Cursor = 0
	}
	p.Clear()
	m.Print(p)
}
func (m *Menu) Prev(p *Printing) {
	if m.Cursor > 0 {
		m.Cursor--
	} else {
		m.Cursor = len(m.Commands) - 1
	}
	p.Clear()
	m.Print(p)
}

func NewMenu() *Menu {
	tmpcs := []Command{}
	i := 1
	for i < 5 {
		tmpc := Command{
			Title: fmt.Sprintf("Command %d", i),
			Cli:   "foo",
		}
		tmpcs = append(tmpcs, tmpc)
		i++
	}
	return &Menu{Commands: tmpcs}
}

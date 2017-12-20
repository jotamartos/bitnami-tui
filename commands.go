package main

import (
	"fmt"
)

type Command struct {
	Title    string
	Cli      string
	Optional bool
	Selected bool
}

type Menu struct {
	Commands      []Command
	Cursor        int
	BottomBar     bool
	BottomBarText string
}

func (m *Menu) SelectToggle() {
	if m.Cursor < len(m.Commands) {
		m.Commands[m.Cursor].Selected = !m.Commands[m.Cursor].Selected
	}
}

func (m *Menu) Print(p *Printing) {
	p.Clear()
	for i, c := range m.Commands {
		title := c.Title
		if c.Optional {
			check := "[ ]"
			if c.Selected {
				check = "[x]"
			}
			title = check + " " + title
		}
		p.Putln(title, i == m.Cursor)
	}
	if m.BottomBar {
		p.BottomBar(m.BottomBarText)
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
			Title:    fmt.Sprintf("Command %d", i),
			Cli:      "foo",
			Optional: true,
		}
		tmpcs = append(tmpcs, tmpc)
		i++
	}
	return &Menu{Commands: tmpcs, BottomBar: true, BottomBarText: "Press ESC to exit"}
}

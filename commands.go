package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
)

type Command struct {
	Title    string
	Cli      string
	Optional bool
	Selected bool
}

type Menu struct {
	Title         string
	Description   string
	Commands      []Command
	Cursor        int
	BottomBar     bool
	BottomBarText string
	Wait          chan int
	p             *Printing
}

func (m *Menu) SelectToggle() {
	if m.Cursor < len(m.Commands) {
		m.Commands[m.Cursor].Selected = !m.Commands[m.Cursor].Selected
	}
}

func (m *Menu) Show() {
	m.p.Clear()
	if m.Title != "" {
		m.p.Putln(m.Title, true)
		m.p.Putln("\n\n", false)
	}

	if m.Description != "" {
		m.p.Putln(m.Description, false)
		m.p.Putln("\n", false)
	}
	for i, c := range m.Commands {
		title := c.Title
		if c.Optional {
			check := "[ ]"
			if c.Selected {
				check = "[x]"
			}
			title = check + " " + title
		}
		m.p.Putln(title, i == m.Cursor)
	}
	if m.BottomBar {
		m.p.BottomBar(m.BottomBarText)
	}
	m.p.Show()
}
func (m *Menu) Quit() {
	m.p.s.Fini()
}

func (m *Menu) Next() {
	if m.Cursor < len(m.Commands)-1 {
		m.Cursor++
	} else {
		m.Cursor = 0
	}
	m.p.Clear()
	m.Show()
}
func (m *Menu) Prev() {
	if m.Cursor > 0 {
		m.Cursor--
	} else {
		m.Cursor = len(m.Commands) - 1
	}
	m.p.Clear()
	m.Show()
}

func NewMenu(style *Style) *Menu {
	channel := make(chan int)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	p := NewPrinting(s, style)
	return &Menu{BottomBar: true, BottomBarText: "Press ESC to exit", Wait: channel, p: p}
}

func (menu *Menu) EventManager() {
	for {
		ev := menu.p.Screen().PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				close(menu.Wait)
				return
			case tcell.KeyEnter:
				menu.SelectToggle()
				menu.Show()
				menu.p.Show()

			case tcell.KeyCtrlL:
				menu.p.Sync()
			case tcell.KeyUp:
				menu.Prev()
				menu.p.Show()
			case tcell.KeyDown:
				menu.Next()
				menu.p.Show()
			}
		case *tcell.EventResize:
			menu.p.Sync()
		}
	}
}

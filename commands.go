package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"
)

type Command struct {
	Title       string
	Description string
	Cli         string
	Execute     HandlerCommand
	Args        []Argument
	Optional    bool
	Selected    bool
}

type Argument struct {
	Envar       string
	Flag        string
	Title       string
	Description string
	Value       string
	IsBoolean   bool
	Valuebool   bool
}

type Menu struct {
	Title         string
	Description   string
	Commands      []Command
	Cursor        int
	BottomBar     bool
	BottomBarText string
	BackText      string
	Wait          chan int
	p             *Printing
}

type HandlerCommand func(arg []Argument)

func (m *Menu) SelectToggle() {
	if m.Cursor < len(m.Commands) {
		m.Commands[m.Cursor].Selected = !m.Commands[m.Cursor].Selected
	}
}
func (m *Menu) IsToggle() bool {
	if m.Cursor < len(m.Commands) {
		return m.Commands[m.Cursor].Optional
	}
	return false
}

func (m *Menu) printPageHearder(title string, desc string) {
	if title != "" {
		m.p.Putln(title, true)
		m.p.Putln("\n\n", false)
	}

	if desc != "" {
		m.p.Putln(desc, false)
		m.p.Putln("\n", false)
	}
}

func (m *Menu) ShowCommand() {
	m.p.Clear()
	if m.Cursor >= len(m.Commands) {
		fmt.Println("Cursor exceed array")
		return
	}
	c := m.Commands[m.Cursor]
	m.printPageHearder(c.Title, c.Description)
	if m.BottomBar {
		m.p.BottomBar(m.BackText)
	}
	m.p.Show()
}

func (m *Menu) Show() {
	m.p.Clear()
	m.printPageHearder(m.Title, m.Description)
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
	return &Menu{BottomBar: true, BottomBarText: "Press ESC to exit", BackText: "Press ESC to go back", Wait: channel, p: p}
}

func (menu *Menu) EventCommandManager() {
	for {
		ev := menu.p.Screen().PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				menu.Show()
				go menu.EventManager()
				return
			case tcell.KeyEnter:
				if menu.IsToggle() {
					menu.SelectToggle()
				}
				menu.ShowCommand()

			case tcell.KeyCtrlL:
				menu.p.Sync()
			}
		case *tcell.EventResize:
			menu.p.Sync()
		}
	}
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
				if menu.IsToggle() {
					menu.SelectToggle()
				}
				menu.ShowCommand()
				go menu.EventCommandManager()
				return

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

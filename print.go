package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
	"os"
)

var styleBitnamiText = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color17).Bold(true)
var styleBitnamiTextHighlight = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.Color17).Bold(true)
var styleBitnamiMenu = tcell.StyleDefault.Background(tcell.ColorSilver).Foreground(tcell.Color17).Bold(true)

func NewPrinting(s tcell.Screen) *Printing {
	encoding.Register()

	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.Color17))
	s.Clear()

	return &Printing{
		s:          s,
		Hightlight: styleBitnamiTextHighlight,
		Default:    styleBitnamiText,
		Menu:       styleBitnamiMenu,
		Indent:     2,
	}
}

type Printing struct {
	s          tcell.Screen
	Cursor     int
	Indent     int
	Hightlight tcell.Style
	Default    tcell.Style
	Menu       tcell.Style
}

func (p *Printing) Clear() {
	p.s.Clear()
	p.Top()
}

func (p *Printing) Screen() tcell.Screen {
	return p.s
}

func (p *Printing) Show() {
	p.s.Show()
}

func (p *Printing) Sync() {
	p.s.Sync()
}

func (p *Printing) Top() {
	p.Cursor = 0
}
func (p *Printing) Bottom() {
	_, p.Cursor = p.s.Size()
	p.Cursor--
}

func (p *Printing) Putln(str string, highlight bool) {
	if highlight {
		p.puts(p.Hightlight, p.Indent, p.Cursor, str)
	} else {
		p.puts(p.Default, p.Indent, p.Cursor, str)
	}
	p.Cursor++
}

func (p *Printing) BottomBar(str string) {
	_, y := p.s.Size()

	p.puts(p.Menu, 0, y-1, "  "+str)
}

func (p *Printing) puts(style tcell.Style, x, y int, str string) {
	i := 0
	var deferred []rune
	dwidth := 0
	for _, r := range str {
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				p.s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				p.s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		p.s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
	xScreen, _ := p.s.Size()
	for i < xScreen {
		p.s.SetContent(x+i, y, ' ', nil, style)
		i++
	}

}

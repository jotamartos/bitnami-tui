package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"github.com/mattn/go-runewidth"
	"os"
)

var row = 0
var indent = 2
var styleBitnamiText = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color17).Bold(true)
var styleBitnamiTextHighlight = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.Color17).Bold(true)

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
		S:          s,
		Hightlight: styleBitnamiTextHighlight,
		Default:    styleBitnamiText,
	}
}

type Printing struct {
	S          tcell.Screen
	Cursor     int
	Indent     int
	Hightlight tcell.Style
	Default    tcell.Style
}

func (p *Printing) Clear() {
	p.S.Clear()
	p.Top()
}

func (p *Printing) Show() {
	p.S.Show()
}

func (p *Printing) Top() {
	p.Cursor = 0
}
func (p *Printing) Bottom() {
	_, p.Cursor = p.S.Size()
	p.Cursor--
}

func (p *Printing) Putln(str string, highlight bool) {
	if highlight {
		puts(p.S, p.Hightlight, p.Indent, p.Cursor, str)
	} else {
		puts(p.S, p.Default, p.Indent, p.Cursor, str)
	}
	p.Cursor++
}

func puts(s tcell.Screen, style tcell.Style, x, y int, str string) {
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
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				s.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		s.SetContent(x+i, y, deferred[0], deferred[1:], style)
		i += dwidth
	}
	xScreen, _ := s.Size()
	for i < xScreen {
		s.SetContent(x+i, y, ' ', nil, style)
		i++
	}

}

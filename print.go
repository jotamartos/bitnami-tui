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

type Style struct {
	Indent     int
	Hightlight tcell.Style
	Default    tcell.Style
	Menu       tcell.Style
	H1         tcell.Style
}

func DefaultStyle() *Style {
	return &Style{
		Hightlight: styleBitnamiTextHighlight,
		Default:    styleBitnamiText,
		Menu:       styleBitnamiMenu,
		H1:         styleBitnamiTextHighlight,
		Indent:     2,
	}
}

func NewPrinting(s tcell.Screen, style *Style) *Printing {
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
		s:     s,
		style: style,
	}
}

type Printing struct {
	s       tcell.Screen
	Cursor  int
	xcursor int
	style   *Style
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
func (p *Printing) Return() {
	p.Cursor++
}

func (p *Printing) Putln(str string, highlight bool) {
	if highlight {
		p.Cursor = p.puts(p.style.Hightlight, p.style.Indent, p.Cursor, str)
	} else {
		p.Cursor = p.puts(p.style.Default, p.style.Indent, p.Cursor, str)
	}
	p.Cursor++
}
func (p *Printing) Put(str string, highlight bool) {
	if highlight {
		p.Cursor = p.puts(p.style.Hightlight, p.style.Indent+p.xcursor, p.Cursor, str)
	} else {
		p.Cursor = p.puts(p.style.Default, p.style.Indent+p.xcursor, p.Cursor, str)
	}
}

func (p *Printing) PutH1(str string, highlight bool) {
	p.Cursor = p.puts(p.style.H1, p.style.Indent, p.Cursor, str)
	p.Cursor++
}

func (p *Printing) BottomBar(str string) {
	_, y := p.s.Size()
	p.puts(p.style.Menu, 0, y-1, "  "+str)
}
func (p *Printing) putc(style tcell.Style, x, y int, str string) int {
	i := 0
	var deferred []rune
	dwidth := 0
	xScreen, _ := p.s.Size()

	for _, r := range str {
		if x+i >= xScreen-p.style.Indent {
			i = 0
			y++
		}
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
	return y
}

func (p *Printing) puts(style tcell.Style, x, y int, str string) int {
	i := 0
	var deferred []rune
	dwidth := 0
	xScreen, _ := p.s.Size()

	for _, r := range str {
		if x+i >= xScreen-p.style.Indent {
			i = 0
			y++
		}
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
	for i < xScreen {
		p.s.SetContent(x+i, y, ' ', nil, style)
		i++
	}
	return y

}

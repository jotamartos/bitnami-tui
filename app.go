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
	m := NewMenu(DefaultStyle())
	m.Title = "Test"
	m.Description = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas id congue felis,vitae auctor metus. Morbi placerat lectus a velit feugiat, ac tincidunt ex ultricies. Nullam fermentum vestibulum tellus, gravida lacinia dui fringilla eget. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.`
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

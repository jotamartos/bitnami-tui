package main

import (
  "fmt"
  "os"
  "encoding/json"
  "github.com/go-ini/ini"
  "github.com/jotamartos/tui"
)

var PROP_FILE   = "/opt/bitnami/properties.ini"
var TUI_FILE   = "/opt/bitnami/btui.json"

const (
  GENERAL         = "General"
  BASE_STACK_KEY  = "base_stack_key"
  BASE_STACK_NAME = "base_stack_name"
  SUPPORT         = "https://community.bitnami.com/"
)

type Stack struct {
	Name string
	Key  string
}

func LoadStack(file string) *Stack {
	cfg, err := ini.Load(file)
	if err != nil {
		fmt.Println("could not find properties file", file)
		return nil
	}
	sec1, err := cfg.GetSection(GENERAL)
	if err != nil {
		fmt.Println("error parsing ini file", err)
		return nil
	}
	keyStack, err := sec1.GetKey(BASE_STACK_KEY)
	if err != nil {
		fmt.Println("error parsing base stack", err)
		return nil
	}
	nameStack, err := sec1.GetKey(BASE_STACK_NAME)
	if err != nil {
		fmt.Println("error parsing base stack name", err)
		return nil
	}

	return &Stack{
		Name: nameStack.Value(),
		Key:  keyStack.Value(),
	}

}

func printMainMenu(stack *Stack, file string) *tui.Menu {
	m := tui.NewMenu(tui.DefaultStyle())
	m.Title = fmt.Sprintf("%s Frequently Run Commands", stack.Name)
	m.Description = `Welcome to Bitnami's frequently run commands, please select from the list below what activities you would like to perform`

  // Open commands.json file to create the menu
  jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("could not find commands.json file", file)
		return nil
	}
  // Close json file and parse it after that
  defer jsonFile.Close()
  decoder := json.NewDecoder(jsonFile)
  tmpcs := []tui.Option{}
  jotaerr := decoder.Decode(&tmpcs)
  if jotaerr != nil {
    fmt.Println("error:", jotaerr)
  }
	m.Options = tmpcs
	return m
}

func main() {
  if _, err := os.Stat(PROP_FILE); os.IsNotExist(err) {
    PROP_FILE     = "./properties.ini"
  }
  if _, err := os.Stat(TUI_FILE); os.IsNotExist(err) {
    TUI_FILE     = "./btui.json"
  }
	stack := LoadStack(PROP_FILE)
	if stack == nil {
		return
	}

	menu := printMainMenu(stack, TUI_FILE)

	menu.PrintMenu()
	go menu.EventManager()
	<-menu.Wait
	menu.Quit()
}

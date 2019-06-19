package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/jotamartos/tui"
)

var (
	// Version of the tool
	version  = ""
	propFile = "/opt/bitnami/properties.ini"
	tuiFile  = "/opt/bitnami/btui/btui.json"
)

const (
	// general states the section in the properties.ini file
	general = "General"
	// baseStackKey states the key related to the stack key in the properties.ini file
	baseStackKey = "base_stack_key"
	// baseStackName states the key related to the stack name in the properties.ini file
	baseStackName = "base_stack_name"
	// supportLink
	supportLink = "https://community.bitnami.com/"
)

// Stack represents an bitnami stack
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
	sec1, err := cfg.GetSection(general)
	if err != nil {
		fmt.Println("error parsing ini file", err)
		return nil
	}
	keyStack, err := sec1.GetKey(baseStackKey)
	if err != nil {
		fmt.Println("error parsing base stack", err)
		return nil
	}
	nameStack, err := sec1.GetKey(baseStackName)
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
	m.Description = fmt.Sprintf(`Welcome to Bitnami's frequently run commands tool (%s), please select from the list below what activities you would like to perform`, version)

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
	if _, err := os.Stat(propFile); os.IsNotExist(err) {
		propFile = "./properties.ini"
	}
	if _, err := os.Stat(tuiFile); os.IsNotExist(err) {
		tuiFile = "./btui.json"
	}
	stack := LoadStack(propFile)
	if stack == nil {
		return
	}

	menu := printMainMenu(stack, tuiFile)

	menu.PrintMenu()
	go menu.EventManager()
	<-menu.Wait
	menu.Quit()
}

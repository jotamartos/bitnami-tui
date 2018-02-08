package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/vtuson/tui"
)

const (
	PROP_FILE       = "/opt/bitnami/properties.ini"
	GENERAL         = "General"
	BASE_STACK      = "base_stack_key"
	BASE_STACK_NAME = "base_stack_name"
	SUPPORT         = "https://community.bitnami.com/"
)

type Stack struct {
	Name string
	Key  string
}

func LoadStack(file string) *Stack {
	cfg, err := ini.Load(PROP_FILE)
	if err != nil {
		fmt.Println("could not find properties file", PROP_FILE)
		return nil
	}
	sec1, err := cfg.GetSection(GENERAL)
	if err != nil {
		fmt.Println("error parsing ini file", err)
		return nil
	}
	keyStack, err := sec1.GetKey(BASE_STACK)
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

func NewTestMenu(stack *Stack) *tui.Menu {
	m := tui.NewMenu(tui.DefaultStyle())
	m.Title = fmt.Sprintf("%s Frequently Run Commands", stack.Name)
	m.Description = `Welcome to Bitnami's frequently run commands, please select from the list below what activities you would like to perform`

	tmpcs := []tui.Command{
		tui.Command{
			Title:       "Remove the Bitnami Banner",
			Cli:         fmt.Sprintf("/opt/bitnami/apps/%s/bnconfig --disable_banner 1", stack.Key),
			Description: "Removing the bitnami banner",
			Success:     "The banner has been removed, if it is still there please go to " + SUPPORT,
			Fail:        "Something when wrong while removing the banner, Please run bnsupport and open a ticket at:" + SUPPORT,
		},
		tui.Command{
			Title:       "Check you service status",
			Cli:         "/opt/bitnami/ctlscript.sh status",
			Description: "Checking service status",
			Success:     "This is your status:",
			Fail:        "Something when wrong while removing the banner, Please run bnsupport and open a ticket at:" + SUPPORT,
			PrintOut:    true,
		},
		tui.Command{
			Title:       "Set up Let's Encrypt",
			Cli:         "/opt/bitnami/letsencrypt/scripts/generate-certificate.sh",
			Description: "Connect your domain with lets encrypt",
			Args: []tui.Argument{
				tui.Argument{
					Description: "Please enter an email associated with your domain",
					Title:       "Your email",
					Name:        "m",
					IsFlag:      true,
				},
				tui.Argument{
					Description: "Please enter your domain name (mydomain.com)",
					Title:       "Your domain",
					Name:        "d",
					IsFlag:      true,
				},
			},
			Success: "SSL via Let's Encrypt is now setup in your application",
			Fail:    "Oh, it didnt work. Please run bnsupport from this tool and contact " + SUPPORT,
		},
		tui.Command{
			Title:       "Set up your Domain for " + stack.Name,
			Cli:         fmt.Sprintf("/opt/bitnami/apps/%s/bnconfig", stack.Key),
			Description: "Connect your domain with " + stack.Name,
			Args: []tui.Argument{
				tui.Argument{
					Description: "Please enter your domain name (mydomain.com)",
					Title:       "Your domain",
					Name:        "-machine_hostname",
					IsFlag:      true,
				},
			},
			Success: "You domain has been set up",
			Fail:    "Oh, it didnt work. Please run bnsupport from this tool and contact " + SUPPORT,
		},
		tui.Command{
			Title:       "Run our support tool (bnsupport)",
			Cli:         "/opt/bitnami/bnsupport-linux-x64.run",
			Description: "Collecting data and uploading results",
			Success:     "Please attached the following ID to your support ticket at " + SUPPORT,
			Fail:        "Something when wrong while removing the banner, Please go to:" + SUPPORT,
			PrintOut:    true,
		},
	}
	m.Commands = tmpcs
	return m
}

func main() {
	stack := LoadStack(PROP_FILE)
	if stack == nil {
		return
	}

	menu := NewTestMenu(stack)

	menu.Show()
	go menu.EventManager()
	<-menu.Wait
	menu.Quit()
}

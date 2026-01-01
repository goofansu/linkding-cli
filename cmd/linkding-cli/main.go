package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/goofansu/linkding-cli/internal/app"
)

type Options struct {
	Login  LoginCommand  `command:"login" description:"Authenticate with Linkding instance"`
	Logout LogoutCommand `command:"logout" description:"Remove stored credentials"`
	Add    AddCommand    `command:"add" description:"Add a new bookmark"`
}

type LoginCommand struct {
	Endpoint string `long:"endpoint" description:"Linkding endpoint URL" required:"yes"`
	APIKey   string `long:"api-key" description:"API Token from Linkding Settings" required:"yes"`
}

type LogoutCommand struct{}

type AddCommand struct {
	Args struct {
		URL string `positional-arg-name:"url" description:"URL of the bookmark to add" required:"yes"`
	} `positional-args:"yes"`
	Notes string `long:"notes" description:"Optional notes for the bookmark (e.g., --notes \"Interesting article\")"`
	Tags  string `long:"tags" description:"Optional tags separated by spaces (e.g., --tags \"golang api\")"`
}

func (c *LoginCommand) Execute(_ []string) error {
	return app.Login(c.Endpoint, c.APIKey)
}

func (c *LogoutCommand) Execute(_ []string) error {
	return app.Logout()
}

func (c *AddCommand) Execute(args []string) error {
	url := c.Args.URL
	return app.AddBookmark(url, c.Notes, c.Tags)
}

func (c *LoginCommand) Usage() string {
	return "[OPTIONS]"
}

func (c *AddCommand) Usage() string {
	return "<url>"
}

func main() {
	opts := Options{}
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	parser.ShortDescription = "Linkding command-line interface"
	parser.LongDescription = "linkding-cli provides a simple CLI for managing Linkding bookmarks.\n\nAuthenticate with login, then use add to add bookmarks.\n\nExamples:\n  linkding-cli login --endpoint https://linkding.example.com --api-key YOUR_API_KEY\n  linkding-cli add https://example.com\n  linkding-cli add https://example.com --notes \"Interesting article\"\n  linkding-cli add https://example.com --tags \"golang api\"\n  linkding-cli logout"

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		return
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				fmt.Fprint(os.Stdout, flagsErr.Message)
				return
			}
			fmt.Fprintf(os.Stderr, "error: %s\n\n", flagsErr.Message)
			parser.WriteHelp(os.Stderr)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"

	"github.com/bscott/googl"
	"github.com/codegangsta/cli"
)

// Commands for this CLI
var Commands = []cli.Command{
	shorten,
	expand,
}

var shorten = cli.Command{
	Name:  "shorten",
	Usage: "Shorten an url",
	Description: `
		Shorten a given url with Google URL Shortener service.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "k, key",
			Usage: "your Google Url Shortener Public API Key",
		},
	},
	Action: Shorten,
}

var expand = cli.Command{
	Name:  "expand",
	Usage: "Expand an url",
	Description: `
		Expand a given short url with Google URL Shortener service.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "k, key",
			Usage: "your Google Url Shortener Public API Key",
		},
	},
	Action: Expand,
}

// Shorten a long URL.
func Shorten(c *cli.Context) {
	url := c.Args().First()
	key := c.String("key")
	client := googl.NewClient(key)
	shortURL, err := client.Shorten(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("shorten Short url:", shortURL.ID)
	fmt.Println("shorten Long url:", shortURL.LongURL)
}

// Expand returns a Long URL.
func Expand(c *cli.Context) {
	url := c.Args().First()
	key := c.String("key")
	client := googl.NewClient(key)
	longURL := client.Expand(url)
	fmt.Println("expand:", longURL)
}

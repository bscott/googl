package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	shorten,
	expand,
}

var shorten = cli.Command{
	Name: "shorten",
	Usage: "Shorten an url",
	Description: `
	Shorten a given url with Google URL Shortener service.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "k, key",
			Usage: "you Google Url Shortener Public API Key",
		},
	},
	Action: Shorten,
}

var expand = cli.Command{
	Name: "expand",
	Usage: "Expand an url",
	Description: `
	Expand a given short url with Google URL Shortener service.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "k, key",
			Usage: "you Google Url Shortener Public API Key",
		},
	},
	Action: Expand,
}

func Shorten(c *cli.Context) {
	url := c.Args().First()
	key := c.String("key")
	fmt.Println("Url:", url)
	fmt.Println("API Key:", key)
}

func Expand(c *cli.Context) {
	url := c.Args().First()
	key := c.String("key")
	fmt.Println("Url:", url)
	fmt.Println("API Key:", key)
}

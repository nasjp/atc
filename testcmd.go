package atc

import (
	"fmt"

	"github.com/urfave/cli"
)

func GetTestCmd() cli.Command {
	return cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test io",
		Action: func(c *cli.Context) error {
			fmt.Println("atc run test")
			return nil
		},
	}
}

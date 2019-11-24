package atc

import (
	"fmt"

	"github.com/NasSilverBullet/atc/internal"
	"github.com/urfave/cli"
)

// GetExampleTestCmd return run exmaple test command
func GetInitCmd() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "initialize config files",
		Action: func(c *cli.Context) error {
			if ok := internal.ExistConfig(); ok {
				fmt.Println("config file aleady exists")
				return nil
			}

			if err := internal.GenerateConfig(); err != nil {
				return err
			}
			fmt.Println("generated config file successfully!!")

			return nil
		},
	}
}

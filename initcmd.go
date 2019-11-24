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
				fmt.Printf("Config file (%s) aleady exists\n", internal.ConfigPath)
				return nil
			}

			if err := internal.GenerateConfig(); err != nil {
				return err
			}
			fmt.Println("Config file generated successfully!!")
			fmt.Printf("Please edit config file (%s)\n", internal.ConfigPath)

			return nil
		},
	}
}

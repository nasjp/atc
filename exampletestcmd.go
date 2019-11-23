package atc

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NasSilverBullet/atc/internal"
	"github.com/urfave/cli"
)

// GetExampleTestCmd return run exmaple test command
func GetExampleTestCmd() cli.Command {
	return cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test io",
		Action: func(c *cli.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			fs, err := ioutil.ReadDir(dir)

			if err != nil {
				return err
			}
			for _, f := range fs {
				fmt.Println(f.Name())
			}
			e, err := internal.GetExamples("abc144", "a")
			if err != nil {
				return err
			}

			_ = e

			return nil
		},
	}
}

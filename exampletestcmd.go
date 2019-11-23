package atc

import (
	"fmt"

	"github.com/NasSilverBullet/atc/internal/task"
	"github.com/urfave/cli"
)

func GetExampleTestCmd() cli.Command {
	return cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test io",
		Action: func(c *cli.Context) error {
			e, err := task.GetExample("abc144", "a")
			if err != nil {
				return err
			}
			fmt.Println(e)
			return nil
		},
	}
}

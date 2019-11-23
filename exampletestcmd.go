package atc

import (
	"github.com/NasSilverBullet/atc/internal/task"
	"github.com/urfave/cli"
)

func GetExampleTestCmd() cli.Command {
	return cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test io",
		Action: func(c *cli.Context) error {
			e, err := task.GetExamples("abc144", "a")
			if err != nil {
				return err
			}
			_ = e
			return nil
		},
	}
}

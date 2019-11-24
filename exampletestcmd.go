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
				c, err := internal.NewConfig()
				if err != nil {
					return err
				}

				lp := c.LanguagesConfig.GetLanguagePath(f.Name())
				if lp == "" {
					continue
				}

				es, err := internal.GetExamples(f.Name())
				if err != nil {
					return err
				}

				if len(es) == 0 {
					continue
				}
				fmt.Printf("=== RUN   %s\n", f.Name())
				for i, e := range es {
					// TODO 実行時間を表示する
					out, err := e.Run(lp, f.Name())
					if err != nil {
						fmt.Printf("    --- FAIL   %s  sample%d\n", f.Name(), i+1)
						fmt.Printf("        Runtime error: %v\n", err)
						continue
					}
					if out == e.Output {
						fmt.Printf("    --- PASS   %s  sample%d\n", f.Name(), i+1)
						continue
					}
					fmt.Printf("    --- FAIL   %s  sample%d\n", f.Name(), i+1)
					fmt.Printf("        Expect: %s", out)
					fmt.Printf("        Actual: %s", e.Output)
				}

			}
			return nil
		},
	}
}

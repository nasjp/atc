package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/atc"
	"github.com/urfave/cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)

}

func run() error {
	app := cli.NewApp()

	app.Name = "atc"
	app.Usage = "light weight test tool for atcoder"
	app.Version = "0.0.1"

	routes(app)

	app.Run(os.Args)
	return nil
}

func routes(app *cli.App) *cli.App {
	app.Commands = []cli.Command{
		atc.GetExampleTestCmd(),
	}
	return app
}

package app

import (
	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	flags := []cli.Flag{
		cli.IntFlag{
			Name:  "start",
			Value: 1,
		},
		cli.IntFlag{
			Name:     "end",
			Required: true,
		},
		cli.StringFlag{
			Name:  "ignore",
			Value: "",
		},
		cli.BoolFlag{
			Name: "back",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "print",
			Flags:  flags,
			Action: calculatePagesOrderToPrint,
		},
	}

	return app
}

package main

import (
	"os"
	"printer-script-go/app"
)

func main() {
	application := app.Generate()

	application.Run(os.Args)
}

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "zaim-cli"
	app.Usage = "zaim-cli"
	app.Action = Action
	app.Run(os.Args)
}

func Action(c *cli.Context) {
	fmt.Println("main")
}

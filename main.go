package main

import (
	"fmt"
	"os"

	"github.com/prog-lang/brainf-go/cpu"
	"github.com/prog-lang/brainf-go/parse"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := app().Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func app() *cli.App {
	return &cli.App{
		Name:   "brainf-go",
		Usage:  "run brainf*ck programs",
		Action: run,
	}
}

func run(c *cli.Context) (err error) {
	name := c.Args().Get(0)
	code, err := parse.FromFile(name)
	if err != nil {
		return
	}
	vm := cpu.Default(code)
	vm.Start()
	if e := vm.Error(); e != nil {
		err = fmt.Errorf("panic during execution: %v", e)
	}
	return nil
}

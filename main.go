package main

import (
	"fmt"
	"os"

	"github.com/prog-lang/brainf-go/cpu"
	"github.com/prog-lang/brainf-go/parse"
	"github.com/urfave/cli/v2"
)

const flagREPL = "repl"

var app = &cli.App{
	Name:           "brainf-go",
	Usage:          "Elegant Brainf*ck Interpreter",
	DefaultCommand: "help",
	Commands: []*cli.Command{{
		Name:      "run",
		Usage:     "Runs program from source code file",
		ArgsUsage: "path/to/source.bf",
		Action:    run,
	}, {
		Name:      "repl",
		Usage:     "Enters the Read, Evaluate, Print, Loop mode",
		ArgsUsage: " ", //! No arguments.
		Action:    repl,
	}},
}

func run(c *cli.Context) error {
	name := c.Args().First()
	code, err := parse.FromFile(name)
	if err != nil {
		return fmt.Errorf("failed to parse source code: %s", err)
	}
	if err := cpu.Default(code).Start().Error(); err != nil {
		return fmt.Errorf("panic during execution: %v", err)
	}
	return nil
}

func repl(_ *cli.Context) error {
	NewREPL().Start()
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

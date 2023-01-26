package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	"github.com/prog-lang/brainf-go/cpu"
	"github.com/prog-lang/brainf-go/parse"
)

type REPL struct {
	*stdio
	notFirstLoop bool
}

type stdio struct {
	*bufio.Reader
	*bufio.Writer
}

func NewREPL() *REPL {
	return &REPL{
		stdio: &stdio{
			Reader: bufio.NewReader(os.Stdin),
			Writer: bufio.NewWriter(os.Stdout),
		},
	}
}

func (repl *REPL) Start() {
	for {
		repl.loop()
	}
}

func (repl *REPL) loop() {
	src, err := repl.read()
	if err != nil {
		log.Printf("failed to read source code: %s", err)
		return
	}
	code, err := parse.FromReader(bytes.NewReader(src))
	if err != nil {
		log.Printf("failed to parse source code: %s", err)
		return
	}
	if err := cpu.Default(code).Start().Error(); err != nil {
		log.Printf("panic during execution: %v", err)
	}
}

func (repl *REPL) read() (src []byte, err error) {
	repl.promptStart()
	src, err = io.ReadAll(repl)
	repl.promptEnd()
	return
}

func (repl *REPL) promptStart() {
	if repl.notFirstLoop {
		repl.WriteByte('\n')
	}
	repl.notFirstLoop = true
	repl.WriteString("❯❯❯\n")
	repl.Flush()
}

func (repl *REPL) promptEnd() {
	repl.WriteString("❮❮❮\n")
	repl.Flush()
}

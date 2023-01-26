# Brainf\*ck Interpreter in Go

Simple to read, well-tested, and elegant Brainf\*ck interpreter written in Go.

## Index

- [parser + compiler](./parse)
- [virtual machine](./cpu)
- [some examples](./examples)

## Install

```bash
go install github.com/prog-lang/brainf-go@latest
```

## Usage

### Help

```text
NAME:
   brainf-go - Elegant Brainf*ck Interpreter

USAGE:
   brainf-go [global options] command [command options] [arguments...]

COMMANDS:
   run      Runs program from source code file
   repl     Enters the Read, Evaluate, Print, Loop mode
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### Example

```bash
brainf-go help
brainf-go repl
brainf-go run ./examples/hello-world.bf
```

**NOTE:** REPL stops reading at EOF. Once you're done inputting the source code,
you can begin code execution with `ENTER > CTRL + D`.

# Brainf*ck Interpreter in Go

Simple to read, well-tested, and elegant Brainf*ck interpreter written in Go.

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
   brainf-go - Elegant Brainf*uck Interpreter

USAGE:
   brainf-go [global options] command [command options] [arguments...]

COMMANDS:
   run, r   Runs program from source code file
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### Example

```bash
brainf-go help
brainf-go run ./examples/hello-world.bf
```

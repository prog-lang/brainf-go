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
   brainf-go - run brainf*ck programs

USAGE:
   brainf-go [global options] [source file name]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### Example

```bash
brainf-go help
brainf-go ./examples/hello-world.bf
```

package parse

import (
	"fmt"

	"github.com/prog-lang/brainf-go/cpu"
	"github.com/zyedidia/generic/stack"
)

type parser struct {
	brackets *stack.Stack[uint]
	ip       uint
	ops      []cpu.Op
}

func newParser() *parser {
	return &parser{
		brackets: stack.New[uint](),
	}
}

func (p *parser) feed(b byte) error {
	switch b {
	case '-':
		p.ops = append(p.ops, cpu.DEC)
	case '+':
		p.ops = append(p.ops, cpu.INC)
	case '<':
		p.ops = append(p.ops, cpu.PREV)
	case '>':
		p.ops = append(p.ops, cpu.NEXT)
	case ',':
		p.ops = append(p.ops, cpu.IN)
	case '.':
		p.ops = append(p.ops, cpu.OUT)
	case '[':
		p.brackets.Push(p.ip)
		p.ops = append(p.ops, cpu.NOP)
	case ']':
		if p.brackets.Size() == 0 {
			return fmt.Errorf("`]` at %d does not have a matching `[`", p.ip)
		}
		openBracketIndex := p.brackets.Pop()
		p.ops[openBracketIndex] = cpu.FWD(p.ip)
		p.ops = append(p.ops, cpu.BACK(openBracketIndex))
	default: // Skip all non-command bytes without incrementing IP.
		return nil
	}
	p.ip++
	return nil
}

func (p *parser) code() ([]cpu.Op, error) {
	if p.brackets.Size() != 0 {
		return nil,
			fmt.Errorf("found %d `[` without a match", p.brackets.Size())
	}
	return p.ops, nil
}

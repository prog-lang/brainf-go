package cpu

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCellOverflow(t *testing.T) {
	code := make([]Op, 0, 300)
	for i := 0; i < 300; i++ {
		code = append(code, INC)
	}
	vm := Default(code)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	const expect = 44
	if cell := vm.tape[vm.tp]; cell != expect {
		t.Fatalf("failed to properly overflow cell value: wanted %d; got %d",
			expect, cell)
	}
}

func TestStdIn(t *testing.T) {
	code := []Op{IN}
	const letter = 'a'
	vm := New(code, bytes.NewReader([]byte{letter}), nil)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	if cell := vm.tape[vm.tp]; cell != letter {
		t.Fatalf("failed to read from stdin: wanted %d; got %d", letter, cell)
	}
}

func TestStdOut(t *testing.T) {
	code := []Op{INC, INC, INC, INC, INC, OUT}
	buf := bytes.NewBuffer([]byte{})
	vm := New(code, nil, buf)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	expect := []byte{5}
	bytes := buf.Bytes()
	if !cmp.Equal(expect, bytes) {
		t.Fatalf("invalid bytes written to buffer:\n%s",
			cmp.Diff(expect, bytes))
	}
}

func TestConditionalJumps(t *testing.T) {
	code := []Op{INC, INC, INC, FWD(6), OUT, DEC, BACK(3)}
	buf := bytes.NewBuffer([]byte{})
	vm := New(code, nil, buf)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	expect := []byte{3, 2, 1}
	bytes := buf.Bytes()
	if !cmp.Equal(expect, bytes) {
		t.Fatalf("invalid bytes written to buffer:\n%s",
			cmp.Diff(expect, bytes))
	}
}

func TestTapeOverflow(t *testing.T) {
	code := []Op{NEXT, NEXT, INC}
	vm := NewWithTapeLength(code, nil, nil, 3)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
}

func TestTapeUnderflow(t *testing.T) {
	code := []Op{PREV, INC, INC, PREV, INC}
	vm := NewWithTapeLength(code, nil, nil, 3)
	vm.Start()
	if err := vm.Error(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	if vm.tape[vm.tp] != 1 {
		t.Fatalf("invalid index move: wanted %d; got %d", 1, vm.tape[vm.tp])
	}
}

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
	if err := vm.Start(); err != nil {
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
	if err := vm.Start(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	if cell := vm.tape[vm.tp]; cell != letter {
		t.Fatalf("failed to read from stdin: wanted %v; got %v", letter, cell)
	}
}

func TestStdOut(t *testing.T) {
	code := []Op{INC, INC, INC, INC, INC, OUT}
	buf := bytes.NewBuffer([]byte{})
	vm := New(code, nil, buf)
	if err := vm.Start(); err != nil {
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
	if err := vm.Start(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	expect := []byte{3, 2, 1}
	bytes := buf.Bytes()
	if !cmp.Equal(expect, bytes) {
		t.Fatalf("invalid bytes written to buffer:\n%s",
			cmp.Diff(expect, bytes))
	}
}

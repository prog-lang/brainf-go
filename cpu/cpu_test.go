package cpu

import "testing"

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

package parse

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	. "github.com/prog-lang/brainf-go/cpu"
)

func TestParser(t *testing.T) {
	source := "+++ [ > +++ [ .- ] < - ]"
	code, err := FromReader(strings.NewReader(source))
	if err != nil {
		t.Fatalf("unexpected parsing error: %s", err)
	}
	buf := bytes.NewBuffer([]byte{})
	vm := New(code, nil, buf)
	if err := vm.Start(); err != nil {
		t.Fatalf("unexpected execution error: %s", err)
	}
	expect := []byte{3, 2, 1, 3, 2, 1, 3, 2, 1}
	got := buf.Bytes()
	if !cmp.Equal(expect, got) {
		t.Fatalf("invalid bytes written to buffer:\n%s", cmp.Diff(expect, got))
	}
}

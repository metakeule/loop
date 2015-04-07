package spinner

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSpin(t *testing.T) {

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	sp := New()

	sp.Spin()
	sp.Spin()
	sp.Spin()
	sp.Spin()
	sp.Spin()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	expected := "\b-\b\\\b|\b/\b-"
	if out != expected {
		t.Errorf("got: %#v; expected: %#v", out, expected)
	}
}

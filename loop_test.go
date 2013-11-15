package loop

import (
	"testing"
)

func TestLoop(t *testing.T) {
	testcases := map[int]string{
		4:  "abca",
		10: "abcabcabca",
	}
	loop := New([]byte("abc"))

	for l, expected := range testcases {
		loop.Reset()
		var all []byte
		for i := 0; i < l; i += 2 {
			get := make([]byte, 2)
			n, err := loop.Read(get)
			if err != nil {
				t.Errorf("loop should never err, but has error: %s\n", err)
			}
			if n != 2 {
				t.Errorf("loop always read length of input, this should be 2 but is: %d\n", n)
			}
			all = append(all, get...)

		}
		if string(all) != expected {
			t.Errorf("wrong result, should be %s, but is: %s\n", expected, string(all))
		}
	}
}

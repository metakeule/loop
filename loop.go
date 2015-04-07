package loop

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Loop struct {
	Data    []byte
	pointer int
	length  int
}

func New(data []byte) *Loop {
	return &Loop{
		Data:    data,
		pointer: 0,
		length:  len(data),
	}
}

func (l *Loop) Reset() { l.pointer = 0 }

// return bytes as long as you are asked for it
func (l *Loop) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		dataPos := (l.pointer + i) % l.length
		p[i] = l.Data[dataPos]
	}
	l.pointer = (l.pointer + len(p)) % l.length
	return len(p), nil
}

type Spinner struct {
	io.Writer
	l *Loop
}

func NewSpinner(w io.Writer) *Spinner {
	return &Spinner{w, New([]byte("-\\|/"))}
}

func (s *Spinner) Spin() (err error) {
	b := make([]byte, 1)
	s.l.Read(b)

	_, err = fmt.Fprint(s.Writer, "\b"+string(b))
	return
}

func (s *Spinner) SpinEvery(d time.Duration) {
	for {
		s.Spin()
		time.Sleep(d)
	}
}

func NewStdoutSpinner() *Spinner {
	return NewSpinner(os.Stdout)
}

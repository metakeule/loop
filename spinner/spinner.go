package spinner

import (
	"gopkg.in/metakeule/loop.v4"
	"os"
	"time"
)

type Spinner struct {
	l *loop.Loop
}

func New() *Spinner {
	return &Spinner{loop.New([]byte("-\\|/"))}
}

func (s *Spinner) Spin() (err error) {
	b := make([]byte, 1)
	s.l.Read(b)

	_, err = os.Stdout.WriteString("\b" + string(b))
	return
}

func (s *Spinner) SpinEvery(d time.Duration) {
	for {
		s.Spin()
		time.Sleep(d)
	}
}

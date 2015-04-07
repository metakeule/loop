package sound

import (
	"gopkg.in/metakeule/loop.v4"
	"os/exec"
)

func PlayAiff(sound []byte) error {
	cmd := exec.Command("aplay", "-f", "cdr")
	cmd.Stdin = loop.New(sound)
	return cmd.Run()
}

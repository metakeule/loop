package main

import (
	"gopkg.in/metakeule/loop.v4/spinner"
	"time"
)

func main() {
	var sp *spinner.Spinner
	fw := spinner.New()
	rev := spinner.NewReverse()
	sp = fw
	for i := 0; i < 1000; i++ {
		sp.Spin()
		time.Sleep(120 * time.Millisecond)
		if i%40 == 0 {
			if sp == fw {
				sp = rev
			} else {
				sp = fw
			}
		}
	}
}

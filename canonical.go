// +build !dev

package loop 	// import "gopkg.in/metakeule/loop.v1"

import (
	"github.com/metakeule/nodouble"
)

const (
	VERSION = "1"
)

func init() {
	nodouble.AddPackage("github.com/metakeule/loop", VERSION)
}


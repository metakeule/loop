// +build !dev

package loop 	// import "gopkg.in/metakeule/loop.v2"

import (
	"github.com/metakeule/nodouble"
)

const (
	VERSION = "2"
)

func init() {
	nodouble.AddPackage("github.com/metakeule/loop", VERSION)
}


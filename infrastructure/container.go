package infrastructure

import (
	"github.com/drgomesp/cargo"
)

type Container cargo.Container

var C *Container

func init() {
	dic := cargo.container.New()

	C = dic
}

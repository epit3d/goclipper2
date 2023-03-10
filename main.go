package main

import (
	"github.com/epit3d/goclipper2/goclipper2"
)

func main() {
	p := goclipper2.NewPaths64()

	p.Simplify()
}

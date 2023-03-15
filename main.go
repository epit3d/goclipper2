package main

import (
	"log"

	"github.com/epit3d/goclipper2/goclipper2"
)

func main() {
	p := goclipper2.NewPath64()

	p.AddPoint(*goclipper2.NewPoint64(0, 0))

	log.Println(p)
}

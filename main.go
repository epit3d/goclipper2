package main

import (
	"log"

	"github.com/epit3d/goclipper2/goclipper2"
)

func main() {
	path := goclipper2.NewCPath64()

	path.AddPoint(goclipper2.NewCPoint64(10, 10))

	log.Println(path)
}

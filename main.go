package main

import (
	"log"

	"github.com/epit3d/goclipper2/goclipper2"
)

func main() {
	co := goclipper2.NewClipperoffset(2.0, 0.0, 0, 0)

	// c64 := goclipper2.NewClipper64()

	input := goclipper2.NewPath64()
	input.AddPoint(*goclipper2.NewPoint64(0, 0, 0))
	input.AddPoint(*goclipper2.NewPoint64(0, 5, 0))
	input.AddPoint(*goclipper2.NewPoint64(5, 5, 0))
	input.AddPoint(*goclipper2.NewPoint64(5, 0, 0))

	co.SetZCallback(func(e1bot, e1top, e2bot, e2top, pt *goclipper2.Point64) {

	})

	co.AddPath64(*input, goclipper2.RoundJoin, goclipper2.PolygonEnd)
	outputs := co.Execute(1)

	log.Println(outputs)
}

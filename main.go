package main

import (
	"log"

	"github.com/epit3d/goclipper2/goclipper2"
)

func main() {
	var cb goclipper2.ClipperOffsetCallback = func(
		path *goclipper2.Path64,
		path_normals *goclipper2.PathD,
		curr_idx int,
		prev_idx int,
	) float64 {
		log.Println("my callback is called with params ", curr_idx, prev_idx)
		return 20.0
	}

	p := goclipper2.NewPath64()

	p.AddPoint(*goclipper2.NewPoint64(0, 0))
	p.AddPoint(*goclipper2.NewPoint64(10, 0))
	p.AddPoint(*goclipper2.NewPoint64(10, 10))
	p.AddPoint(*goclipper2.NewPoint64(0, 10))
	p.AddPoint(*goclipper2.NewPoint64(0, 0))

	co := goclipper2.NewClipperoffset(0.2, 0, 0, 0)
	co.AddPath64(*p, goclipper2.MiterJoin, goclipper2.PolygonEnd)
	log.Println("check me")

	result := co.ExecuteCallback(cb)
	log.Println(result)
}

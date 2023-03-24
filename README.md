# goclipper2

### golang bindings for [Clipper2](https://github.com/AngusJohnson/Clipper2) - a polygon <a href="https://en.wikipedia.org/wiki/Clipping_(computer_graphics)">Clipping</a> and <a href="https://en.wikipedia.org/wiki/Parallel_curve">Offsetting</a> library (originally in C++, C# &amp; Delphi)<br>

### Documentation:

I would gladly refer you to original [documentation](http://www.angusj.com/clipper2/Docs/Overview.htm) by **AngusJohnson** and ask to search for similar names.

### Usage

Get library:

```bash
go get github.com/epit3d/goclipper2
```

Simple code:

```go
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

```

```bash
go run main.go
```

If you do this outside this repo, you should copy `lib` directory to the root of application.

### For developer:

If you want to dig into how everything is building, go [here](READMEdev.md)

### Contribution

Feel free to create an issue or PR with your thoughts.

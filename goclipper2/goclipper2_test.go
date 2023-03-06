package goclipper2_test

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/epit3d/goclipper2/goclipper2"
	"github.com/stretchr/testify/require"
)

func IsNonPath(line string) bool {
	return len(line) == 0 || (!unicode.IsDigit(rune(line[0])) && line[0] != '-')
}

type test struct {
	ct       goclipper2.ClipperClipType
	fr       goclipper2.ClipperFillRule
	area     int
	count    int
	subj     *goclipper2.Paths64
	subjOpen *goclipper2.Paths64
	clip     *goclipper2.Paths64
}

func ParseTest(s string) (*test, error) {
	res := test{
		subj:     goclipper2.NewPaths64(),
		subjOpen: goclipper2.NewPaths64(),
		clip:     goclipper2.NewPaths64(),
	}

	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.Contains(line, "INTERSECTION") {
			res.ct = goclipper2.Intersection
		} else if strings.Contains(line, "UNION") {
			res.ct = goclipper2.Union
		} else if strings.Contains(line, "DIFFERENCE") {
			res.ct = goclipper2.Difference
		} else if strings.Contains(line, "XOR") {
			res.ct = goclipper2.XOR
		} else if strings.Contains(line, "EVENODD") {
			res.fr = goclipper2.EvenOdd
		} else if strings.Contains(line, "NONZERO") {
			res.fr = goclipper2.NonZero
		} else if strings.Contains(line, "POSITIVE") {
			res.fr = goclipper2.Positive
		} else if strings.Contains(line, "NEGATIVE") {
			res.fr = goclipper2.Negative
		} else if strings.Contains(line, "SOL_AREA") {
			if _, err := fmt.Sscanf(line, "SOL_AREA: %d", &res.area); err != nil {
				return nil, err
			}
		} else if strings.Contains(line, "SOL_COUNT") {
			if _, err := fmt.Sscanf(line, "SOL_COUNT: %d", &res.count); err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(line, "SUBJECTS_OPEN") {
			// TODO: organize same parsing of paths
			i++
			for i < len(lines) && !IsNonPath(lines[i]) {
				newPath, err := goclipper2.NewPath64OfString(lines[i])
				if err != nil {
					log.Fatalf("do not expect error here, err: <%v>", err)
					return nil, err
				}

				res.subjOpen.AddPath(*newPath)
				i++
			}
			i--
		} else if strings.HasPrefix(line, "SUBJECTS") {
			i++
			for i < len(lines) && !IsNonPath(lines[i]) {
				newPath, err := goclipper2.NewPath64OfString(lines[i])
				if err != nil {
					log.Fatalf("do not expect error here, err: <%v>", err)
					return nil, err
				}

				res.subj.AddPath(*newPath)
				i++
			}
			i--
		} else if strings.HasPrefix(line, "CLIPS") {
			i++
			for i < len(lines) && !IsNonPath(lines[i]) {
				newPath, err := goclipper2.NewPath64OfString(lines[i])
				if err != nil {
					log.Fatalf("do not expect error here, err: <%v>", err)
					return nil, err
				}

				res.clip.AddPath(*newPath)
				i++
			}
			i--
		}
	}

	return &res, nil
}

func TestLines(t *testing.T) {
	// here we have to parse lines
	fname := "../third_party/clipper2c/vendor/Clipper2/Tests/Lines.txt"

	f, err := os.Open(fname)
	require.NoError(t, err)
	defer f.Close()

	fstr, err := io.ReadAll(f)
	require.NoError(t, err)

	tests := strings.Split(string(fstr), "\n\n")
	// log.Println(len(tests))

	for i, testStr := range tests[1:] {
		test, err := ParseTest(testStr)
		require.NoError(t, err)

		c := goclipper2.NewClipper64()
		c.AddSubject(*test.subj)
		c.AddOpenSubject(*test.subjOpen)
		c.AddClip(*test.clip)

		p1 := goclipper2.NewPaths64()
		p2 := goclipper2.NewPaths64()

		result := c.Execute(test.ct, test.fr, *p1, *p2)
		require.Equal(t, result, int64(1))

		count2 := p1.Length() + p2.Length()
		countDiff := int64(math.Abs(float64(count2 - int64(test.count))))

		require.LessOrEqual(t, int64(countDiff), int64(8))
		log.Println("test", i, "passed")
	}
}

func TestOffsetOrientation(t *testing.T) {
	co := goclipper2.NewClipperoffset(2.0, 0.0, 0, 0)

	input := goclipper2.NewPath64()
	input.AddPoint(*goclipper2.NewPoint64(0, 0))
	input.AddPoint(*goclipper2.NewPoint64(0, 5))
	input.AddPoint(*goclipper2.NewPoint64(5, 5))
	input.AddPoint(*goclipper2.NewPoint64(5, 0))

	co.AddPath64(*input, goclipper2.RoundJoin, goclipper2.PolygonEnd)
	outputs := co.Execute(1)

	require.Equal(t, outputs.Length(), int64(1))
	require.True(t, (input.Area() >= 0) == (outputs.GetPath(0).Area() >= 0))
}

func TestOffsets(t *testing.T) {
	// here we have to parse lines
	fname := "../third_party/clipper2c/vendor/Clipper2/Tests/Offsets.txt"

	f, err := os.Open(fname)
	require.NoError(t, err)
	defer f.Close()

	fstr, err := io.ReadAll(f)
	require.NoError(t, err)

	tests := strings.Split(string(fstr), "\n\n")
	log.Println("there are ", len(tests), "tests")

	for i, testStr := range tests {
		test, err := ParseTest(testStr)
		// log.Println(test.subj)
		require.NoError(t, err)

		co := goclipper2.NewClipperoffset(2.0, 0.0, 0, 0)
		co.AddPaths64(*test.subj, goclipper2.RoundJoin, goclipper2.PolygonEnd)
		outputs := co.Execute(1)

		// is the sum total area of the solution is positive
		outerIsPositive := outputs.Area() > 0

		// there should be exactly one exterior path
		isPositiveCount := 0
		isNegativeCount := 0

		for i := 0; i < int(outputs.Length()); i++ {
			if outputs.GetPath(int64(i)).Area() >= 0 {
				isPositiveCount++
			} else {
				isNegativeCount++
			}
		}

		if outerIsPositive {
			require.Equal(t, 1, isPositiveCount)
		} else {
			require.Equal(t, 1, isNegativeCount)
		}

		log.Println("test", i, "passed")
	}
}

func TestNegativeOrientation(t *testing.T) {
	subjects := goclipper2.NewPaths64()
	p, err := goclipper2.NewPath64OfString("0,0, 0,100, 100,100, 100,0")
	require.NoError(t, err)
	subjects.AddPath(*p)
	require.False(t, p.IsPositive() == 1)

	p, err = goclipper2.NewPath64OfString("10,10, 10,110, 110,110, 110,10")
	require.NoError(t, err)
	subjects.AddPath(*p)
	require.False(t, p.IsPositive() == 1)

	clips := goclipper2.NewPaths64()
	p, err = goclipper2.NewPath64OfString("50,50, 50,150, 150,150, 150,50")
	require.NoError(t, err)
	clips.AddPath(*p)

	solution := subjects.Union(*clips, goclipper2.Negative)
	require.Equal(t, solution.Length(), int64(1))
	require.Equal(t, solution.GetPath(0).Length(), int64(12))
}

func TestPolygons(t *testing.T) {
	// here we have to parse lines
	fname := "../third_party/clipper2c/vendor/Clipper2/Tests/Polygons.txt"

	f, err := os.Open(fname)
	require.NoError(t, err)
	defer f.Close()

	fstr, err := io.ReadAll(f)
	require.NoError(t, err)

	tests := strings.Split(string(fstr), "\n\n")
	log.Println("there are ", len(tests), "tests")

	isInList := func(v int, vals []int) bool {
		for _, vi := range vals {
			if vi == v {
				return true
			}
		}

		return false
	}

	for i, testStr := range tests {
		test, err := ParseTest(testStr)
		require.NoError(t, err)

		c := goclipper2.NewClipper64()
		c.AddSubject(*test.subj)
		c.AddOpenSubject(*test.subjOpen)
		c.AddClip(*test.clip)
		var (
			solution     = goclipper2.NewPaths64()
			solutionOpen = goclipper2.NewPaths64()
		)
		c.Execute(test.ct, test.fr, *solution, *solutionOpen)

		// get metrics
		measuredArea := solution.Area()
		areaDiff := math.Abs(measuredArea - float64(test.area))
		areaDiffRatio := math.Abs(float64(areaDiff) / measuredArea)

		measuredCount := solution.Length() + solutionOpen.Length()
		countDiff := int(math.Abs(float64(measuredCount - int64(test.count))))

		// solve the same for polytree
		cTree := goclipper2.NewClipper64()
		cTree.AddSubject(*test.subj)
		cTree.AddOpenSubject(*test.subjOpen)
		cTree.AddClip(*test.clip)
		var (
			solutionTree  = goclipper2.NewPolytree64(goclipper2.PolyTree64{})
			solutionOpen2 = goclipper2.NewPaths64()
			// solutionTreeOpen = goclipper2.NewPolytree64(goclipper2.PolyTree64{})
		)
		cTree.ExecuteTreeWithOpen(test.ct, test.fr, *solutionTree, *solutionOpen2)

		measuredAreaPolytree := solutionTree.Area()
		measuredCountPolytree := solutionTree.ToPaths().Length()

		testNumber := i + 1

		// check polygon counts
		if test.count <= 0 {
		} else if isInList(testNumber, []int{120, 121, 130, 138,
			140, 148, 163, 165, 166, 167, 168, 172, 175, 178, 180}) {
			require.LessOrEqual(t, countDiff, 5)
		} else if isInList(testNumber, []int{27, 181}) {
			require.LessOrEqual(t, countDiff, 2)
		} else if testNumber >= 120 && testNumber <= 184 {
			require.LessOrEqual(t, countDiff, 2)
		} else if isInList(testNumber, []int{23, 45, 87, 102, 111, 113, 191}) {
			require.LessOrEqual(t, countDiff, 1)
		} else {
			require.Equal(t, countDiff, 0)
		}

		// check polygon areas
		if test.area <= 0 {
		} else if isInList(testNumber, []int{19, 22, 23, 24}) {
			require.LessOrEqual(t, areaDiffRatio, 0.5)
		} else if testNumber == 193 {
			require.LessOrEqual(t, areaDiffRatio, 0.2)
		} else if testNumber == 63 {
			require.LessOrEqual(t, areaDiffRatio, 0.1)
		} else if testNumber == 16 {
			require.LessOrEqual(t, areaDiffRatio, 0.075)
		} else if testNumber == 26 {
			require.LessOrEqual(t, areaDiffRatio, 0.05)
		} else if isInList(testNumber, []int{15, 52, 53, 54, 59, 60, 64, 117, 119, 184}) {
			require.LessOrEqual(t, areaDiffRatio, 0.02)
		} else {
			require.LessOrEqual(t, areaDiffRatio, 0.01)
		}

		require.Equal(t, measuredCount, measuredCountPolytree)
		require.Equal(t, measuredArea, measuredAreaPolytree)

		log.Println("test", testNumber, "passed")
	}
}

func TestPolyTreeHoles1(t *testing.T) {
	// here we have to parse lines
	fname := "../third_party/clipper2c/vendor/Clipper2/Tests/PolytreeHoleOwner.txt"

	f, err := os.Open(fname)
	require.NoError(t, err)
	defer f.Close()

	fstr, err := io.ReadAll(f)
	require.NoError(t, err)

	tests := strings.Split(string(fstr), "\n\n")
	log.Println("there are ", len(tests), "tests")
	require.GreaterOrEqual(t, len(tests), 1)
	test, err := ParseTest(tests[0])
	require.NoError(t, err)

	c := goclipper2.NewClipper64()
	c.AddSubject(*test.subj)
	c.AddOpenSubject(*test.subjOpen)
	c.AddClip(*test.clip)
	var (
		solution     = goclipper2.NewPolytree64(goclipper2.PolyTree64{})
		solutionOpen = goclipper2.NewPaths64()
	)

	c.ExecuteTreeWithOpen(test.ct, test.fr, *solution, *solutionOpen)
	require.Equal(t, solution.FullyContainsChildren(), int64(1))
}

func TestPolyTreeIntersection(t *testing.T) {
	c := goclipper2.NewClipper64()

	subject := goclipper2.NewPaths64()
	p, err := goclipper2.NewPath64OfString("0,0, 0,5 5,5, 5,0")
	require.NoError(t, err)
	subject.AddPath(*p)
	c.AddSubject(*subject)

	clip := goclipper2.NewPaths64()
	p, err = goclipper2.NewPath64OfString("1,1, 1,6, 6,6 6,1")
	require.NoError(t, err)
	clip.AddPath(*p)
	c.AddClip(*clip)

	var (
		solution  = goclipper2.NewPolytree64(goclipper2.PolyTree64{})
		openPaths = goclipper2.NewPaths64()
	)

	if subject.GetPath(0).IsPositive() == 1 {
		c.ExecuteTreeWithOpen(goclipper2.Intersection, goclipper2.Positive, *solution, *openPaths)
	} else {
		c.ExecuteTreeWithOpen(goclipper2.Intersection, goclipper2.Negative, *solution, *openPaths)
	}

	require.Equal(t, int64(0), openPaths.Length())
	require.Equal(t, int64(1), solution.Count())
	require.Equal(t, int64(4), solution.GetChild(0).Polygon().Length())
}

func TestPolytreeUnion(t *testing.T) {
	c := goclipper2.NewClipper64()

	subject := goclipper2.NewPaths64()

	p, err := goclipper2.NewPath64OfString("0,0, 0,5 5,5, 5,0")
	require.NoError(t, err)
	subject.AddPath(*p)

	p, err = goclipper2.NewPath64OfString("1,1, 1,6, 6,6 6,1")
	require.NoError(t, err)
	subject.AddPath(*p)

	c.AddSubject(*subject)

	var (
		solution  = goclipper2.NewPolytree64(goclipper2.PolyTree64{})
		openPaths = goclipper2.NewPaths64()
	)

	if subject.GetPath(0).IsPositive() == 1 {
		c.ExecuteTreeWithOpen(goclipper2.Union, goclipper2.Positive, *solution, *openPaths)
	} else {
		// because clipping ops normally return positive solutions
		c.SetReverseSolution(1)
		c.ExecuteTreeWithOpen(goclipper2.Union, goclipper2.Negative, *solution, *openPaths)
	}

	require.Equal(t, int64(0), openPaths.Length())
	require.Equal(t, int64(1), solution.Count())
	require.Equal(t, int64(8), solution.GetChild(0).Polygon().Length())
	require.Equal(t, subject.GetPath(0).IsPositive(), solution.GetChild(0).Polygon().IsPositive())
}

func TestRectClip(t *testing.T) {
	var (
		sub = goclipper2.NewPaths64()
		clp = goclipper2.NewPaths64()
		sol = goclipper2.NewPaths64()
	)

	rect := goclipper2.NewRect64(100, 100, 700, 500)
	clp.AddPath(*rect.AsPath())

	sub.AddPath(*goclipper2.NewPath64MustOfString("100,100, 700,100, 700,500, 100,500"))
	sol = sub.RectClip(*rect)
	require.Equal(t, sol.Area(), sub.Area())
	sub.Delete()

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("110,110, 700,100, 700,500, 100,500"))
	sol = sub.RectClip(*rect)
	require.Equal(t, sol.Area(), sub.Area())
	sub.Delete()

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("90,90, 700,100, 700,500, 100,500"))
	sol = sub.RectClip(*rect)
	require.Equal(t, sol.Area(), clp.Area())
	sub.Delete()

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("90,90, 710,90, 710,510, 90,510"))
	sol = sub.RectClip(*rect)
	require.Equal(t, sol.Area(), clp.Area())
	sub.Delete()

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("110,110, 690,110, 690,490, 110,490"))
	sol = sub.RectClip(*rect)
	require.Equal(t, sol.Area(), sub.Area())
	sub.Delete()

	rect = goclipper2.NewRect64(390, 290, 410, 310)
	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("410,290, 500,290, 500,310, 410,310"))
	sol = sub.RectClip(*rect)
	require.Equal(t, int64(0), sol.Length())

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("430,290, 470,330, 390,330"))
	sol = sub.RectClip(*rect)
	require.Equal(t, int64(0), sol.Length())

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("450,290, 480,330, 450,330"))
	sol = sub.RectClip(*rect)
	require.Equal(t, int64(0), sol.Length())

	sub = goclipper2.NewPaths64()
	sub.AddPath(*goclipper2.NewPath64MustOfString("208,66, 366,112, 402,303," +
		"234,332, 233,262, 243,140, 215,126, 40,172"))
	rect = goclipper2.NewRect64(237, 164, 322, 248)
	sol = sub.RectClip(*rect)
	solBounds := sol.Bounds()
	require.Equal(t, solBounds.Width(), rect.Width())
	require.Equal(t, solBounds.Height(), rect.Height())
}

func TestTrimCollinear(t *testing.T) {
	input1 := goclipper2.NewPath64MustOfString("10,10, 10,10, 50,10, 100,10, 100,100, 10,100, 10,10, 20,10")
	output1 := input1.TrimCollinear(0)
	require.Equal(t, int64(4), output1.Length())

	input2 := goclipper2.NewPath64MustOfString("10,10, 10,10, 100,10, 100,100, 10,100, 10,10, 10,10")
	output2 := input2.TrimCollinear(1)
	require.Equal(t, int64(5), output2.Length())

	input3 := goclipper2.NewPath64MustOfString("10,10, 10,50, 10,10, 50,10,50,50," +
		"50,10, 70,10, 70,50, 70,10, 50,10, 100,10, 100,50, 100,10")
	output3 := input3.TrimCollinear(0)
	require.Equal(t, int64(0), output3.Length())

	input4 := goclipper2.NewPath64MustOfString("2,3, 3,4, 4,4, 4,5, 7,5," +
		"8,4, 8,3, 9,3, 8,3, 7,3, 6,3, 5,3, 4,3, 3,3, 2,3")
	output4a := input4.TrimCollinear(0)
	output4b := output4a.TrimCollinear(0)

	area4a := output4a.Area()
	area4b := output4b.Area()
	require.Equal(t, int64(7), output4a.Length())
	require.Equal(t, -9.0, area4a)
	require.Equal(t, output4a.Length(), output4b.Length())
	require.Equal(t, area4a, area4b)
}

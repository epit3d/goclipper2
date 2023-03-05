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

		result := c.Execute(test.ct, test.fr, p1, p2)
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

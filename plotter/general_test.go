// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
)

var generateTestData = flag.Bool("regen", false, "Uses the current state to regenerate the test data.")

// checkPlot checks a generated plot against a previously created reference.
// If generateTestData = true, it regenerates the reference first.
func checkPlot(dir, name, ext string, p *plot.Plot, width, height vg.Length, handleError errFunc, lf logFunc) {
	filename := filepath.Join(dir, fmt.Sprintf("%s.%s", name, ext))

	c, err := p.WriterTo(width, height, ext)
	handleError(err)

	var buf bytes.Buffer
	_, err = c.WriteTo(&buf)
	handleError(err)

	// Recreate Golden images.
	if *generateTestData {
		handleError(p.Save(width, height, filename))
	}

	f, err := os.Open(filename)
	handleError(err)

	want, err := ioutil.ReadAll(f)
	handleError(err)
	f.Close()
	if !bytes.Equal(buf.Bytes(), want) {
		handleError(fmt.Errorf("image mismatch for %s\n", filename))
		return
	}
	lf("Image can be seen at https://github.com/gonum/plot/tree/master/plotter/%s/%s.%s.\n",
		dir, name, ext)
}

// errFunc is an interface for functions that deal with errors.
type errFunc func(err error)

// handleEx prints err if err!=nil
var handleEx = func(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// handleTest returns a function that fails the test if err!=nil
var handleTest = func(t *testing.T) errFunc {
	return func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
}

type logFunc func(format string, args ...interface{})

var testLog = func(t *testing.T) logFunc {
	return func(s string, arg ...interface{}) {
		t.Logf(s, arg...)
	}
}

var exampleLog = func(s string, arg ...interface{}) {
	fmt.Printf(s, arg...)
}

// Draw the plot logo.
func Example() {
	p, err := plot.New()
	handleEx(err)

	DefaultLineStyle.Width = vg.Points(1)
	DefaultGlyphStyle.Radius = vg.Points(3)

	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
	})
	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
	})

	pts := XYs{{0, 0}, {0, 1}, {0.5, 1}, {0.5, 0.6}, {0, 0.6}}
	line, err := NewLine(pts)
	handleEx(err)
	scatter, err := NewScatter(pts)
	handleEx(err)
	p.Add(line, scatter)

	pts = XYs{{1, 0}, {0.75, 0}, {0.75, 0.75}}
	line, err = NewLine(pts)
	handleEx(err)
	scatter, err = NewScatter(pts)
	handleEx(err)
	p.Add(line, scatter)

	pts = XYs{{0.5, 0.5}, {1, 0.5}}
	line, err = NewLine(pts)
	handleEx(err)
	scatter, err = NewScatter(pts)
	handleEx(err)
	p.Add(line, scatter)

	checkPlot("examplePlots", "plotLogo", "png", p, 100, 100,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/plotLogo.png.
}

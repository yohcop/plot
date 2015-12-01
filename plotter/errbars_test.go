// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"math/rand"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg/draw"
)

// ExampleErrors draws points and error bars.
func ExampleErrors() {

	randomError := func(n int) Errors {
		err := make(Errors, n)
		for i := range err {
			err[i].Low = rand.Float64()
			err[i].High = rand.Float64()
		}
		return err
	}
	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) XYs {
		pts := make(XYs, n)
		for i := range pts {
			if i == 0 {
				pts[i].X = rand.Float64()
			} else {
				pts[i].X = pts[i-1].X + rand.Float64()
			}
			pts[i].Y = pts[i].X + 10*rand.Float64()
		}
		return pts
	}

	type errPoints struct {
		XYs
		YErrors
		XErrors
	}

	rand.Seed(int64(0))
	n := 15
	data := errPoints{
		XYs:     randomPoints(n),
		YErrors: YErrors(randomError(n)),
		XErrors: XErrors(randomError(n)),
	}

	p, err := plot.New()
	handleEx(err)
	scatter, err := NewScatter(data)
	handleEx(err)
	scatter.Shape = draw.CrossGlyph{}
	xerrs, err := NewXErrorBars(data)
	handleEx(err)
	yerrs, err := NewYErrorBars(data)
	handleEx(err)
	p.Add(scatter, xerrs, yerrs)

	checkPlot("examplePlots", "errorBars", "png", p, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/errorBars.png.
	// Normally, you would use plot.Save().
}

func TestErrors(t *testing.T) {
	randomError := func(n int) Errors {
		err := make(Errors, n)
		for i := range err {
			err[i].Low = rand.Float64()
			err[i].High = rand.Float64()
		}
		return err
	}
	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) XYs {
		pts := make(XYs, n)
		for i := range pts {
			if i == 0 {
				pts[i].X = rand.Float64()
			} else {
				pts[i].X = pts[i-1].X + rand.Float64()
			}
			pts[i].Y = pts[i].X + 10*rand.Float64()
		}
		return pts
	}

	type errPoints struct {
		XYs
		YErrors
		XErrors
	}

	rand.Seed(int64(0))
	n := 15
	data := errPoints{
		XYs:     randomPoints(n),
		YErrors: YErrors(randomError(n)),
		XErrors: XErrors(randomError(n)),
	}

	p, err := plot.New()
	handleTest(t)(err)
	scatter, err := NewScatter(data)
	handleTest(t)(err)
	scatter.Shape = draw.CrossGlyph{}
	xerrs, err := NewXErrorBars(data)
	handleTest(t)(err)
	yerrs, err := NewYErrorBars(data)
	handleTest(t)(err)
	p.Add(scatter, xerrs, yerrs)
	p.Add(NewGlyphBoxes())

	checkPlot("examplePlots", "errorBarsTest", "png", p, 200, 200,
		handleTest(t), testLog(t))
}

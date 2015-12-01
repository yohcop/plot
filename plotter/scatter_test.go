// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"math/rand"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

// ExampleScatter draws some scatter points, a line,
// and a line with points.
func ExampleScatter() {
	rand.Seed(int64(0))

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

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	p, err := plot.New()
	handleEx(err)
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(NewGrid())

	s, err := NewScatter(scatterData)
	handleEx(err)
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := NewLine(lineData)
	handleEx(err)
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lpLine, lpPoints, err := NewLinePoints(linePointsData)
	handleEx(err)
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)

	checkPlot("examplePlots", "scatter", "png", p, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/scatter.png.
	// Normally, you would use plot.Save().
}

func TestScatter(t *testing.T) {
	rand.Seed(int64(0))

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

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(NewGrid())

	s, err := NewScatter(scatterData)
	handleTest(t)(err)
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := NewLine(lineData)
	handleTest(t)(err)
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lpLine, lpPoints, err := NewLinePoints(linePointsData)
	handleTest(t)(err)
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)
	// Add a GlyphBox plotter for debugging.
	p.Add(NewGlyphBoxes())

	checkPlot("examplePlots", "scatterTest", "png", p, 200, 200,
		handleTest(t), testLog(t))

}

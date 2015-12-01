// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"math"
	"math/rand"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg"
)

// An example of making a histogram.
func ExampleHistogram() {
	// stdNorm returns the probability of drawing a
	// value from a standard normal distribution.
	stdNorm := func(x float64) float64 {
		const sigma = 1.0
		const mu = 0.0
		const root2π = 2.50662827459517818309
		return 1.0 / (sigma * root2π) * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
	}

	rand.Seed(int64(0))
	n := 10000
	vals := make(Values, n)
	for i := 0; i < n; i++ {
		vals[i] = rand.NormFloat64()
	}

	p, err := plot.New()
	handleEx(err)
	p.Title.Text = "Histogram"
	h, err := NewHist(vals, 16)
	handleEx(err)
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := NewFunction(stdNorm)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	checkPlot("examplePlots", "histogram", "png", p, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/tree/master/plotter/examplePlots/histogram.png.
	// Normally, you would use plot.Save().
}

func TestHistogram(t *testing.T) {
	// stdNorm returns the probability of drawing a
	// value from a standard normal distribution.
	stdNorm := func(x float64) float64 {
		const sigma = 1.0
		const mu = 0.0
		const root2π = 2.50662827459517818309
		return 1.0 / (sigma * root2π) * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
	}

	rand.Seed(int64(0))
	n := 10000
	vals := make(Values, n)
	for i := 0; i < n; i++ {
		vals[i] = rand.NormFloat64()
	}

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Histogram"
	h, err := NewHist(vals, 16)
	handleTest(t)(err)
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := NewFunction(stdNorm)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)
	p.Add(NewGlyphBoxes())

	checkPlot("examplePlots", "histogramTest", "png", p, 200, 200,
		handleTest(t), testLog(t))
}

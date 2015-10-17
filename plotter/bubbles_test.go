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
)

func TestBubblesRadius(t *testing.T) {
	b := &Bubbles{
		MinRadius: vg.Length(0),
		MaxRadius: vg.Length(1),
	}

	tests := []struct {
		minz, maxz, z float64
		r             vg.Length
	}{
		{0, 0, 0, vg.Length(0.5)},
		{1, 1, 1, vg.Length(0.5)},
		{0, 1, 0, vg.Length(0)},
		{0, 1, 1, vg.Length(1)},
		{0, 1, 0.5, vg.Length(0.5)},
		{0, 2, 1, vg.Length(0.5)},
		{0, 4, 0, vg.Length(0)},
		{0, 4, 1, vg.Length(0.25)},
		{0, 4, 2, vg.Length(0.5)},
		{0, 4, 3, vg.Length(0.75)},
		{0, 4, 4, vg.Length(1)},
	}

	for _, test := range tests {
		b.MinZ, b.MaxZ = test.minz, test.maxz
		if r := b.radius(test.z); r != test.r {
			t.Errorf("Got incorrect radius (%g) on %v", r, test)
		}
	}
}

func ExampleBubbles() {

	// randomTriples returns some random x, y, z triples
	// with some interesting kind of trend.
	randomTriples := func(n int) XYZs {
		data := make(XYZs, n)
		for i := range data {
			if i == 0 {
				data[i].X = rand.Float64()
			} else {
				data[i].X = data[i-1].X + 2*rand.Float64()
			}
			data[i].Y = data[i].X + 10*rand.Float64()
			data[i].Z = data[i].X
		}
		return data
	}

	rand.Seed(int64(0))
	n := 10
	bubbleData := randomTriples(n)

	p, err := plot.New()
	handleEx(err)
	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := NewBubbles(bubbleData, vg.Points(1), vg.Points(20))
	handleEx(err)
	bs.Color = color.RGBA{R: 196, B: 128, A: 255}
	p.Add(bs)

	checkPlot("examplePlots", "bubbles", "png", p, 200, 200,
		handleEx, exampleLog)

	// Output:
	// Image can be seen at https://github.com/gonum/plot/plotter/examplePlots/bubbles.png.
	// Normally, you would use plot.Save().
}

func TestBubbles(t *testing.T) {

	// randomTriples returns some random x, y, z triples
	// with some interesting kind of trend.
	randomTriples := func(n int) XYZs {
		data := make(XYZs, n)
		for i := range data {
			if i == 0 {
				data[i].X = rand.Float64()
			} else {
				data[i].X = data[i-1].X + 2*rand.Float64()
			}
			data[i].Y = data[i].X + 10*rand.Float64()
			data[i].Z = data[i].X
		}
		return data
	}

	rand.Seed(int64(0))
	n := 10
	bubbleData := randomTriples(n)

	p, err := plot.New()
	handleTest(t)(err)
	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := NewBubbles(bubbleData, vg.Points(1), vg.Points(20))
	handleTest(t)(err)
	bs.Color = color.RGBA{R: 196, B: 128, A: 255}
	p.Add(bs)
	p.Add(NewGlyphBoxes())

	checkPlot("examplePlots", "bubblesTest", "png", p, 200, 200,
		handleTest(t), testLog(t))
}

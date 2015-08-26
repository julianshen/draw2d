// Copyright 2015 The draw2d Authors. All rights reserved.
// created: 26/06/2015 by Stani Michiels
// See also test_test.go

package draw2dpdf_test

import (
	"testing"

	"github.com/julianshen/draw2d"
	"github.com/julianshen/draw2d/samples/android"
	"github.com/julianshen/draw2d/samples/frameimage"
	"github.com/julianshen/draw2d/samples/geometry"
	"github.com/julianshen/draw2d/samples/gopher"
	"github.com/julianshen/draw2d/samples/gopher2"
	"github.com/julianshen/draw2d/samples/helloworld"
	"github.com/julianshen/draw2d/samples/line"
	"github.com/julianshen/draw2d/samples/linecapjoin"
	"github.com/julianshen/draw2d/samples/postscript"
)

func TestSampleAndroid(t *testing.T) {
	test(t, android.Main)
}

// TODO: FillString: w (width) is incorrect
func TestSampleGeometry(t *testing.T) {
	// Set the global folder for searching fonts
	// The pdf backend needs for every ttf file its corresponding
	// json/.z file which is generated by gofpdf/makefont.
	draw2d.SetFontFolder("../resource/font")
	test(t, geometry.Main)
}

func TestSampleGopher(t *testing.T) {
	test(t, gopher.Main)
}

func TestSampleGopher2(t *testing.T) {
	test(t, gopher2.Main)
}

func TestSampleHelloWorld(t *testing.T) {
	// Set the global folder for searching fonts
	// The pdf backend needs for every ttf file its corresponding
	// json/.z file which is generated by gofpdf/makefont.
	draw2d.SetFontFolder("../resource/font")
	test(t, helloworld.Main)
}

func TestSampleFrameImage(t *testing.T) {
	test(t, frameimage.Main)
}

func TestSampleLine(t *testing.T) {
	test(t, line.Main)
}

func TestSampleLineCap(t *testing.T) {
	test(t, linecapjoin.Main)
}

func TestSamplePostscript(t *testing.T) {
	test(t, postscript.Main)
}

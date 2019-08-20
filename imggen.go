package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
)

// Image is..
type Image struct {
	pixels [][]uint8
	width  int
	height int
	model  color.Model
}

func (i Image) ColorModel() color.Model {
	return i.model
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := i.pixels[y][x]

	return color.RGBA{v, v, 255, 255}
}

func (i *Image) generate(fn func(int, int) uint8) {
	for h := 0; h < i.height; h++ {
		i.pixels[h] = make([]uint8, i.width)

		for w := 0; w < i.width; w++ {
			i.pixels[h][w] = fn(h, w)
		}
	}
}

func test1(x, y int) uint8 {
	return uint8(math.Sqrt(float64(x * y)))
}

func blue(x, y int) uint8 {
	return uint8(0)
}

func makeimage(x, y int) Image {
	// Generates an image by default
	m := Image{make([][]uint8, y), x, y, color.RGBAModel}
	m.generate(blue)

	return m
}

func writeimage(m image.Image, fn string) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fn, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

}

func main() {
	m := makeimage(255, 255)
	writeimage(m, "white.png")
	m.generate(test1)
	writeimage(m, "img.png")
}

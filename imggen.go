package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
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
	return uint8(x + y)
}

func makeimage(x, y int) Image {
	m := Image{make([][]uint8, y), x, y, color.RGBAModel}
	m.generate(test1)

	return m
}

func showimage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

func main() {
	m := makeimage(255, 255)
	showimage(m)
}

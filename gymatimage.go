package gymat

import (
	"image"
	"image/color"
)

type GmImage struct {
	*Gm
}

func (g *GmImage) At(x int, y int) color.Color {
	return color.Gray{g.Gm.Data[y][x]}
}
func (g *GmImage) ColorModel() color.Model {
	return color.GrayModel
}
func (g *GmImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{g.Gm.X, g.Gm.Y},
	}
}

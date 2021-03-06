package colony

import (
	"image"
	"image/color"
	"image/draw"

	"math"
	"math/rand"

	"engo.io/engo/common"
)

type DustRegion struct {
	Region *Region
}

func (dr *DustRegion) GenerateTiles() {
	dr.Region.MakeTiles()

	for i, _ := range dr.Region.Tiles {
		var tile *Tile

		if rand.Float32() > 0.5 {
			tile = dustpatch()
		} else {
			tile = gravel()
		}

		dr.Region.Tiles[i] = tile
	}
}

func (dr *DustRegion) Drawable(size float32) (common.Drawable, error) {
	black := color.NRGBA{A: 255}

	img := uniformimg(black, size, size)

	const g = 200
	gray := color.NRGBA{R: g, G: g, B: g, A: 255}

	const margin = 2
	isz := int(math.Floor(float64(size)))

	rect := image.Rectangle{
		Min: image.Point{X: margin, Y: margin},
		Max: image.Point{X: isz - margin, Y: isz - margin},
	}

	draw.Draw(img, rect, image.NewUniform(gray), image.ZP, draw.Src)

	return imgtexture(img), nil
}

func (dr *DustRegion) ShortName() string {
	return "Dust"
}

type GravelTile struct {
	Tile *Tile
}

func gravel() *Tile {
	t := &Tile{}
	t.Class = &GravelTile{Tile: t}

	return t
}

func (gt *GravelTile) Init() {
}

func (gt *GravelTile) ShortName() string {
	return "Gravel"
}

func init() {
	const gravelseed = 10

	_gravelimg = noiseimg(gravelseed, _gravelsize)
}

const _gravelsize = 128
var _gravelimg *image.NRGBA
func (gt *GravelTile) Drawable(size float32) (common.Drawable, error) {
	if size > _gravelsize {
		panic("Gravel size exceeds maximum")
	}
	isz := int(math.Floor(float64(size)))
	rect := image.Rectangle{Max: image.Point{X: isz, Y: isz}}
	img := image.NewNRGBA(rect)

	draw.Draw(img, rect, _gravelimg, image.ZP, draw.Src)

	return imgtexture(img), nil
}

type DustTile struct {
	Tile *Tile
}

func dustpatch() *Tile {
	t := &Tile{}
	t.Class = &DustTile{Tile: t}

	return t
}

func (dt *DustTile) Init() {
}

func (dt *DustTile) ShortName() string {
	return "Dust"
}

func (dt *DustTile) Drawable(size float32) (common.Drawable, error) {
	const gr = 200
	gray := color.NRGBA{R: gr, G: gr, B: gr, A: 255}

	img := uniformimg(gray, size, size)

	return imgtexture(img), nil
}

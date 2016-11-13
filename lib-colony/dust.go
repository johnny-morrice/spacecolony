package colony

import (
	"image"
	"image/color"
	"image/draw"

	"math"
	"math/rand"

	"github.com/ojrac/opensimplex-go"

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
	const gray = 200
	c := color.NRGBA{R: gray, G: gray, B: gray, A: 255}

	img := uniformimg(c, size, size)

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
	const sz = 128
	const gravelseed = 10

	rect := image.Rectangle{Max: image.Point{X:sz, Y: sz}}
	_gravelimg = image.NewNRGBA(rect)

	noise := opensimplex.NewWithSeed(gravelseed)

	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			shade := noise.Eval2(float64(i), float64(j))
			gr := uint8(math.Floor(shade + 1) * 125)
			col := color.NRGBA{R: gr, G: gr, B: gr, A: 255}

			_gravelimg.SetNRGBA(i, j, col)
		}
	}
}

var _gravelimg *image.NRGBA
func (gt *GravelTile) Drawable(size float32) (common.Drawable, error) {
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

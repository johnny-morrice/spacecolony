package colony

import (
	"image/color"

	"engo.io/engo/common"
)

type DustRegion struct {
	Region *Region
}

func (dr *DustRegion) GenerateTiles(rand *Random) {
	for i, _ := range dr.Region.Tiles {
		dr.Region.Tiles[i] = dr.dustpatch()
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

func (dr *DustRegion) dustpatch() *Tile {
	t := &Tile{}
	t.Class = &DustTile{Tile: t}

	return t
}

type DustTile struct {
	Tile *Tile
}

func (dt *DustTile) Generate(rand *Random) {
	dt.Tile.Type = TileDust
}

func (dt *DustTile) ShortName() string {
	return "Dust"
}

func (dt *DustTile) Drawable(size float32) (common.Drawable, error) {
	return nil, nil
}

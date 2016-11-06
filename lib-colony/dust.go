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

func (dr *DustRegion) Texture(size float64) (*common.Texture, error) {
	fnt := stdfont()
	fnt.Size = size

	const gray = 200
	fnt.FG = color.NRGBA{R: gray, G: gray, B: gray, A: 255}

	err := fnt.CreatePreloaded()

	if err != nil {
		return nil, err
	}

	texture := fnt.Render("#")

	return &texture, nil
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

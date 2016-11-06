package colony

import (
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

func (dr *DustRegion) Texture() (*common.Texture, error) {
	return common.LoadedSprite("sprite/grey-dust.png")
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

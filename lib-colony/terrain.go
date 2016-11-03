package colony

import (
	"fmt"

	"engo.io/engo/common"
)

type Planet struct {
        Tiles []*Region

        Width int
        Height int
}

func (p *Planet) Init(rand *Random) {
	p.Tiles = make([]*Region, p.Width * p.Height)

	for i, _ := range p.Tiles {
		p.Tiles[i] = RandomRegion(rand)
	}
}

type RegionClass interface {
	GenerateTiles(rand *Random)
	Drawable() (common.Drawable, error)
}

type Region struct {
        Biome *Biome

	Class RegionClass

        Tiles []*Tile

        Width int
        Height int

        Neighbours []*Region
}

func RandomRegion(rand *Random) *Region {
	r := &Region{}

	r.Biome = RandomBiome(rand)
	r.Width = DefaultRegionWidth
	r.Height = DefaultRegionHeight

	r.Init(rand)

	return r
}

func (r *Region) Init(rand *Random) {
	r.Tiles = make([]*Tile, r.Width * r.Height)

	switch r.Biome.Type {
	case BiomeDust:
		r.Class = &DustRegion{Region: r}
	default:
		panic(fmt.Sprintf("Unknown BiomeType: %v", r.Biome.Type))
	}

	r.Class.GenerateTiles(rand)
}

type Biome struct {
        Type BiomeType
        Shape BiomeShapeType
}

func RandomBiome(rand *Random) *Biome {
        flatDust := &Biome{}
	flatDust.Type = BiomeDust
	flatDust.Shape = BiomeShapeFlat

	return flatDust
}

type BiomeShapeType uint8

const (
        BiomeShapeFlat = BiomeShapeType(iota)
        BiomeShapeCrater
        BiomeShapeHill
        BiomeShapeMountain
)

type BiomeType uint16

const (
        BiomeDust = BiomeType(iota)
)

const DefaultRegionWidth = 1000
const DefaultRegionHeight = 1000

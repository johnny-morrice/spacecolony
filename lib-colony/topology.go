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
	Drawable(size float32) (common.Drawable, error)
	ShortName() string
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
	switch r.Biome.Type {
	case BiomeDust:
		r.Class = &DustRegion{Region: r}
	default:
		panic(fmt.Sprintf("Unknown BiomeType: %v", r.Biome.Type))
	}

	// r.Tiles = make([]*Tile, r.Width * r.Height)

	// r.Class.GenerateTiles(rand)
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

type Tile struct {
	Type TileType

	Class TileClass
}

type TileType uint8

type TileClass interface {
	Generate(rand *Random)
	ShortName() string
	Drawable(size float32) (common.Drawable, error)
}

const (
	TileDust = TileType(iota)
)

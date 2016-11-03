package colony

import (
	"fmt"
)

type Planet struct {
        Tiles []*Region

        Width uint
        Height uint
}

func (p *Planet) Init(rand Random) {
}

type RegionClass interface {
	GenerateTiles(rand *Random)
}

type Region struct {
        Biome *Biome

	Class RegionClass

        Tiles []*Tile

        Width uint
        Height uint

        Neighbours []*Region
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

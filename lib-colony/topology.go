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

func (p *Planet) Init() {
	p.Tiles = make([]*Region, p.Width * p.Height)

	for i, _ := range p.Tiles {
		p.Tiles[i] = RandomRegion()
	}
}

type RegionClass interface {
	GenerateTiles()
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

func RandomRegion() *Region {
	r := &Region{}

	r.Biome = RandomBiome()
	r.Width = DefaultRegionWidth
	r.Height = DefaultRegionHeight

	r.Init()

	return r
}

func (r *Region) Init() {
	switch r.Biome.Type {
	case BiomeDust:
		r.Class = &DustRegion{Region: r}
	default:
		panic(fmt.Sprintf("Unknown BiomeType: %v", r.Biome.Type))
	}
}

func (r *Region) MakeTiles() {
	r.Tiles = make([]*Tile, r.Width * r.Height)
}

type Biome struct {
        Type BiomeType
        Shape BiomeShapeType
}

func RandomBiome() *Biome {
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

// TODO Tile used to have more members, now not sure about the structure.
type Tile struct {
	Class TileClass
}


type TileClass interface {
	Init()
	ShortName() string
	Drawable(size float32) (common.Drawable, error)
}

// TODO much larger crushes the CPU: implement sparse matrix.
const DefaultRegionWidth = 100
const DefaultRegionHeight = 100

package colony

type PlanetMap struct {
        Tiles []Region

        Width uint
        Height uint

        Config PlanetConfig
}

func (pm *PlanetMap) Generate(rand Random, config PlanetConfig) {
}

type PlanetConfig struct {
        // E.g. GravityStrength int
}

type Region struct {
        Bio Biome

        Tiles []UnitTile

        Width uint
        Height uint

        Neighbours []*Region
}

func (pm *Region) TransformTiles(rand Random) {
}

type Biome struct {
        Type BiomeType
        Shape RegionShapeType
}

func RandomBiome(rand Random) *Biome {
        return nil
}

type RegionShapeType uint8

const (
        RegionShapeFlat = RegionShapeType(iota)
        RegionShapeCrater
        RegionShapeHill
        RegionShapeMountain
)

type BiomeType uint16

const (
        BiomeDust = BiomeType(iota)
)

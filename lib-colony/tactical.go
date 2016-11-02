package colony

type Tile struct {
	Type TileType

	Class TileClass
}

type TileType uint8

type TileClass interface {
	Generate(rand *Random)
}

const (
	TileDust = TileType(iota)
)

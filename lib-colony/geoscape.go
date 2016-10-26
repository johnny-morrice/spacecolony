package colony

import (
	"engo.io/ecs"
	// "engo.io/engo"
	"engo.io/engo/common"
)

type GeoscapeSystem struct {
        drawn bool
        world *ecs.World
}

func (geosys *GeoscapeSystem) Update(dt float32) {
        if !geosys.drawn {
                geosys.regen()

                geosys.drawn = true
        }
}

func (geosys *GeoscapeSystem) Remove(ecs.BasicEntity) {
}

func (geosys *GeoscapeSystem) regen() {
}

type GeoTile struct {
        ecs.BasicEntity
        common.RenderComponent
        common.SpaceComponent
}

type Planet struct {
        Tiles []*GeoTile
        Width int
        Height int
}

func (planet *Planet) FillRandom() {
}

func (planet *Planet) TileAt(x, y int) *GeoTile {
        return planet.Tiles[planet.pos(x, y)]
}

func (planet *Planet) SetTile(x, y int, tile *GeoTile) {
        planet.Tiles[planet.pos(x, y)] = tile
}

func (planet *Planet) pos(x, y int) int {
        return (y * planet.Width) + x
}

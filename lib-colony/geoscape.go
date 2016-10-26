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

	Bounds Bounds

        Planet *PlanetMap
}

type ArielView struct {
        Tiles []*GeoTile
        Width int
        Height int
}

func (av *ArielView) Populate(landscape *PlanetMap) {
}

func (av *ArielView) TileAt(x, y int) *GeoTile {
        return av.Tiles[strideindex(x, y, av.Width)]
}

func (av *ArielView) SetTile(x, y int, tile *GeoTile) {
        av.Tiles[strideindex(x, y, av.Width)] = tile
}

package colony

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GeoscapeSystem struct {
	Tilesize float32

        drawn bool

        world *ecs.World

	planet *Planet
}

func (geosys *GeoscapeSystem) New(w *ecs.World) {
	geosys.world = w
}

func (geosys *GeoscapeSystem) Update(dt float32) {
        if !geosys.drawn {
                geosys.regen()

                geosys.drawn = true

		for i := 0; i < geosys.planet.Width; i++ {
			for j := 0; j < geosys.planet.Height; j++ {
				region := geosys.planet.Tiles[strideindex(i, j, geosys.planet.Width)]

				const margin = 2
				fi := float32(i)
				fj := float32(j)
				x := (fi * geosys.Tilesize) + (fi * margin)
				y := (fj * geosys.Tilesize) + (fj * margin)

				geotile := GeoTile{}

				geotile.BasicEntity = ecs.NewBasic()

				geotile.RegionComponent = RegionComponent{Region: region}

				geotile.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{x, y},
					Width: geosys.Tilesize,
					Height: geosys.Tilesize,
				}

				drawable, err := region.Class.Drawable()

				if err != nil {
					panic(err)
				}

				geotile.RenderComponent = common.RenderComponent{
					Drawable: drawable,
					Scale: engo.Point{1, 1},
				}

				for _, system := range geosys.world.Systems() {
					switch sys := system.(type) {
					case *common.RenderSystem:
						sys.Add(&geotile.BasicEntity, &geotile.RenderComponent, &geotile.SpaceComponent)
					}
				}
			}
		}
        }
}

func (geosys *GeoscapeSystem) Remove(ecs.BasicEntity) {
}

func (geosys *GeoscapeSystem) regen() {
	rand := &Random{}

	geosys.planet = &Planet{}
	geosys.planet.Width = 80
	geosys.planet.Height = 80

	geosys.planet.Init(rand)
}

type GeoTile struct {
        ecs.BasicEntity
        common.RenderComponent
        common.SpaceComponent
	RegionComponent
}

type RegionComponent struct {
	Region *Region
}

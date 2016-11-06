package colony

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GeoscapeSystem struct {
	Tilesize float32
	ScreenWidth float32
	ScreenHeight float32

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
				geosys.addtile(i, j)
			}
		}
        }
}

func (geosys *GeoscapeSystem) addtile(i, j int) {
	region := geosys.planet.Tiles[strideindex(i, j, geosys.planet.Width)]

	planetWidth := float32(geosys.planet.Width) * geosys.Tilesize
	planetHeight := float32(geosys.planet.Height) * geosys.Tilesize
	xOffset := (geosys.ScreenWidth - planetWidth) / 2
	yOffset := (geosys.ScreenHeight - planetHeight) / 2

	if xOffset < 0 || yOffset < 0 {
		panic("Small windows not supported")
	}

	fi := float32(i)
	fj := float32(j)
	x := (fi * geosys.Tilesize) + xOffset
	y := (fj * geosys.Tilesize) + yOffset

	geotile := GeoTile{}

	geotile.BasicEntity = ecs.NewBasic()

	geotile.RegionComponent = RegionComponent{Region: region}

	geotile.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{x, y},
		Width: geosys.Tilesize,
		Height: geosys.Tilesize,
	}

	texture, err := region.Class.Texture()

	if err != nil {
		panic(err)
	}

	scale := geosys.Tilesize / texture.Width()

	geotile.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale: engo.Point{scale, scale},
	}

	for _, system := range geosys.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&geotile.BasicEntity, &geotile.RenderComponent, &geotile.SpaceComponent)
		}
	}
}

func (geosys *GeoscapeSystem) Remove(ecs.BasicEntity) {
}

func (geosys *GeoscapeSystem) regen() {
	rand := &Random{}

	geosys.planet = &Planet{}
	geosys.planet.Width = 40
	geosys.planet.Height = 40

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

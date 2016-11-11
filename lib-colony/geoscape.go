package colony

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type geoscapeScene struct {
	DisplayOptions
	EngineOptions
}

func (gs *geoscapeScene) Type() string { return "geoscape" }

func (gs *geoscapeScene) Preload() {
	err := loadAllAssets()

	if err != nil {
		panic(err)
	}
}

func (gs *geoscapeScene) Setup(world *ecs.World) {
        common.SetBackground(color.Black)

        world.AddSystem(&common.RenderSystem{})
        world.AddSystem(&common.MouseSystem{})

	geoscape := &GeoscapeSystem{}
	geoscape.Tilesize = float32(gs.Tilesize)
	geoscape.ScreenWidth = float32(gs.Width)
	geoscape.ScreenHeight = float32(gs.Height)

        world.AddSystem(geoscape)
}

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

		geosys.embarktext()

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
		Position: engo.Point{X: x, Y: y},
		Width: geosys.Tilesize,
		Height: geosys.Tilesize,
	}

	drawable, err := region.Class.Drawable(float64(geosys.Tilesize) * 1.2)

	if err != nil {
		panic(err)
	}

	geotile.RenderComponent = common.RenderComponent{
		Drawable: drawable,
		Scale: engo.Point{X: 1, Y: 1},
	}

	for _, system := range geosys.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&geotile.BasicEntity, &geotile.RenderComponent, &geotile.SpaceComponent)
		}
	}
}

func (geosys *GeoscapeSystem) embarktext() {
	const titleSize = 50
	texture, err := basicText("Select Landing Zone", titleSize)

	if err != nil {
		panic(err)
	}

	const y = 10
	x := (geosys.ScreenWidth - texture.Width()) / 2

	hud := HudSection{}
	hud.BasicEntity = ecs.NewBasic()
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: x, Y: y},
		Width: texture.Width(),
		Height: texture.Height(),
	}

	hud.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale: engo.Point{X: 1, Y: 1},
	}

	for _, system := range geosys.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
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

type HudSection struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
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

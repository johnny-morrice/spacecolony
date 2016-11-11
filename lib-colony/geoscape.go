package colony

import (
	"image/color"
	"math"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type geoscapeScene struct {
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
	geoscape.ScreenWidth = float32(gs.Width)
	geoscape.ScreenHeight = float32(gs.Height)

        world.AddSystem(geoscape)
}

type GeoscapeSystem struct {
	TileSize float32
	GeoSquareSize float32
	OffsetX float32
	OffsetY float32
	ScreenWidth float32
	ScreenHeight float32

        drawn bool

        world *ecs.World

	planet *Planet
}

func (geosys *GeoscapeSystem) New(w *ecs.World) {
	geosys.world = w

	bound := geosys.ScreenWidth
	bigger := geosys.ScreenHeight
	if bound > bigger {
		bigger, bound = bound, bigger
	}

	margin := bound / 4
	geosys.GeoSquareSize = bound - margin
	geosys.OffsetX = (geosys.ScreenWidth - geosys.GeoSquareSize) / 2
	geosys.OffsetY = (geosys.ScreenHeight - geosys.GeoSquareSize) / 2
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

	fi, fj := float32(i), float32(j)
	var margin float32 = 2
	x := (fi * geosys.TileSize) + geosys.OffsetX + (fi * margin)
	y := (fj * geosys.TileSize) + geosys.OffsetY + (fj * margin)

	geotile := GeoTile{}

	geotile.BasicEntity = ecs.NewBasic()

	geotile.RegionComponent = RegionComponent{Region: region}

	geotile.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: x, Y: y},
		Width: geosys.TileSize,
		Height: geosys.TileSize,
	}

	regionsize := geosys.TileSize - margin

	drawable, err := region.Class.Drawable(regionsize)

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
	titleSize := (geosys.ScreenHeight - geosys.GeoSquareSize) / 6
	texture, err := basictext("Select Landing Zone", titleSize)

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

	const planetsize = 40
	geosys.planet = &Planet{}
	geosys.planet.Width = planetsize
	geosys.planet.Height = planetsize
	geosys.TileSize = evenfloor(geosys.GeoSquareSize / planetsize)

	geosys.planet.Init(rand)
}

func evenfloor(x float32) float32 {
	x = float32(math.Floor(float64(x)))

	if int(x) % 2 == 0 {
		return x
	} else {
		return x - 1
	}
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

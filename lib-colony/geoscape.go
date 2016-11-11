package colony

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GeoscapeScene struct {
	ScreenDims
}

func (gs *GeoscapeScene) Type() string { return "geoscape" }

func (gs *GeoscapeScene) Preload() {
	err := loadAllAssets()

	if err != nil {
		panic(err)
	}
}

func (gs *GeoscapeScene) Setup(world *ecs.World) {
        common.SetBackground(color.Black)

        world.AddSystem(&common.RenderSystem{})
        world.AddSystem(&common.MouseSystem{})

	geosys := &GeoscapeSystem{}
	geosys.TileView = NewTileView(gs.ScreenDims)

        world.AddSystem(geosys)
}

type GeoscapeSystem struct {
	TileView

        drawn bool

        world *ecs.World

	planet *Planet

	tiles []*GeoTile
	regioninfo *HudSection
}

func (geosys *GeoscapeSystem) New(w *ecs.World) {
	geosys.world = w
}

func (geosys *GeoscapeSystem) Update(dt float32) {
        if geosys.drawn {
		geosys.updatehud()
	} else {
                geosys.regen()
		geosys.embarktext()

		geosys.drawn = true
        }
}

func (geosys *GeoscapeSystem) updatehud() {
	geosys.wipeinfo()

	for _, geotile := range geosys.tiles {
		if geotile.Hovered {
			geosys.displayinfo(geotile)
			break
		}
	}
}

func (geosys *GeoscapeSystem) wipeinfo() {
	if geosys.regioninfo != nil {
		derenderentity(geosys.world, &geosys.regioninfo.BasicEntity)
	}

	geosys.regioninfo = nil
}

func (geosys *GeoscapeSystem) displayinfo(geotile *GeoTile) {
	size := (geosys.ScreenHeight - geosys.ViewSquareSize) / 12

	position := func(texture *common.Texture) (float32, float32) {
		return (geosys.ScreenWidth - texture.Width()) / 2, geosys.ScreenWidth - 10 - size
	}

	msg := fmt.Sprintf("%v (%v,%v)", geotile.Region.Class.ShortName(), geotile.X, geotile.Y)

	hud := hudmsg(msg, size, position)

	geosys.regioninfo = hud

	renderentity(geosys.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
}

func (geosys *GeoscapeSystem) addtile(i, j int) {
	geotile := &GeoTile{}

	geotile.BasicEntity = ecs.NewBasic()

	region := geosys.planet.Tiles[strideindex(i, j, geosys.planet.Width)]
	geotile.RegionComponent = RegionComponent{X: i, Y: j, Region: region}

	const margin = 2
	regionsize := geosys.TileSize - margin

	drawable, err := region.Class.Drawable(regionsize)

	if err != nil {
		panic(err)
	}

	geotile.RenderComponent = common.RenderComponent{
		Drawable: drawable,
		Scale: engo.Point{X: 1, Y: 1},
	}

	fi, fj := float32(i), float32(j)
	x := (fi * geosys.TileSize) + geosys.OffsetX + (fi * margin)
	y := (fj * geosys.TileSize) + geosys.OffsetY + (fj * margin)

	geotile.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: x, Y: y},
		Width: geosys.TileSize,
		Height: geosys.TileSize,
	}

	geosys.tiles = append(geosys.tiles, geotile)

	mouseentity(geosys.world, &geotile.BasicEntity, &geotile.MouseComponent, &geotile.RenderComponent, &geotile.SpaceComponent)
	renderentity(geosys.world, &geotile.BasicEntity, &geotile.RenderComponent, &geotile.SpaceComponent)
}

func (geosys *GeoscapeSystem) embarktext() {
	titleSize := (geosys.ScreenHeight - geosys.ViewSquareSize) / 6

	position := func(texture *common.Texture) (float32, float32) {
		return (geosys.ScreenWidth - texture.Width()) / 2, 10
	}

	hud := hudmsg("Select Landing Zone", titleSize, position)

	renderentity(geosys.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
}

func (geosys *GeoscapeSystem) Remove(ecs.BasicEntity) {
}

func (geosys *GeoscapeSystem) regen() {
	rand := &Random{}

	const planetsize = 40
	geosys.planet = &Planet{}
	geosys.planet.Width = planetsize
	geosys.planet.Height = planetsize
	geosys.TileSize = evenfloor(geosys.ViewSquareSize / planetsize)

	geosys.planet.Init(rand)

	for i := 0; i < geosys.planet.Width; i++ {
		for j := 0; j < geosys.planet.Height; j++ {
			geosys.addtile(i, j)
		}
	}
}

type GeoTile struct {
        ecs.BasicEntity
        common.RenderComponent
        common.SpaceComponent
	common.MouseComponent

	RegionComponent
}

type RegionComponent struct {
	X, Y int
	Region *Region
}

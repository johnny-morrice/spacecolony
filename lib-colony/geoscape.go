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
	CenterTiles
	TileSize float32

	planet *Planet
	world *ecs.World
}

func (scene *GeoscapeScene) Type() string { return "geoscape" }

func (scene *GeoscapeScene) Preload() {
	err := loadAllAssets()

	if err != nil {
		panic(err)
	}
}

func (scene *GeoscapeScene) Setup(world *ecs.World) {
	common.SetBackground(color.Black)

	scene.CenterTiles = NewCenterTiles(scene.ScreenDims)
	scene.ScreenDims = scene.ScreenDims
	scene.world = world

	const planetsize = 40
	scene.planet = &Planet{}
	scene.planet.Width = planetsize
	scene.planet.Height = planetsize
	scene.TileSize = evenfloor(scene.ViewSquareSize / planetsize)

	lander := &GeoscapeLander{}
	lander.ScreenDims = scene.ScreenDims
	lander.TileSize = scene.TileSize


	scene.planet.Init()

        world.AddSystem(&common.RenderSystem{})
        world.AddSystem(&common.MouseSystem{})
	world.AddSystem(lander)

	scene.gentiles()
	scene.embarktext()

}

func (scene *GeoscapeScene) addtile(i, j int) {
	geotile := &geotile{}

	geotile.BasicEntity = ecs.NewBasic()

	region := scene.planet.Tiles[strideindex(i, j, scene.planet.Width)]
	geotile.RegionComponent = RegionComponent{X: i, Y: j, Region: region}

	const margin = 2
	regionsize := scene.TileSize - margin

	drawable, err := region.Class.Drawable(regionsize)

	if err != nil {
		panic(err)
	}

	geotile.RenderComponent = rndcomp(drawable)

	fi, fj := float32(i), float32(j)
	x := (fi * scene.TileSize) + scene.VSMinX + (fi * margin)
	y := (fj * scene.TileSize) + scene.VSMinY + (fj * margin)

	geotile.SpaceComponent = spacecompsz(x, y, scene.TileSize, scene.TileSize)

	mouseentity(scene.world, &geotile.BasicEntity, &geotile.MouseComponent, &geotile.RenderComponent, &geotile.SpaceComponent)
	renderentity(scene.world, &geotile.BasicEntity, &geotile.RenderComponent, &geotile.SpaceComponent)

	for _, system := range scene.world.Systems() {
		switch sys := system.(type) {
		case *GeoscapeLander:
			sys.Add(&geotile.BasicEntity, &geotile.MouseComponent, &geotile.RegionComponent)
		}
	}
}

func (scene *GeoscapeScene) embarktext() {
	titleSize := (scene.ScreenHeight - scene.ViewSquareSize) / 6

	position := func(texture *common.Texture) (float32, float32) {
		return (scene.ScreenWidth - texture.Width()) / 2, 10
	}

	hud := hudmsg("Select Landing Zone", titleSize, position)

	renderentity(scene.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
}

func (scene *GeoscapeScene) gentiles() {
	for i := 0; i < scene.planet.Width; i++ {
		for j := 0; j < scene.planet.Height; j++ {
			scene.addtile(i, j)
		}
	}
}

type mouseregion struct {
	*common.MouseComponent
	*RegionComponent
}

type GeoscapeLander struct {
	ScreenDims
	TileSize float32
	world *ecs.World

	mrs []*mouseregion
	regioninfo *HudSection
}

func (lander *GeoscapeLander) New(w *ecs.World) {
	lander.world = w
}

func (lander *GeoscapeLander) Update(df float32) {
	lander.wipeinfo()
	lander.hoverinfo()
	lander.chooselanding()
}

func (lander *GeoscapeLander) Add(basic *ecs.BasicEntity, m *common.MouseComponent, r *RegionComponent) {
	mr := &mouseregion{
		MouseComponent: m,
		RegionComponent: r,
	}

	lander.mrs = append(lander.mrs, mr)
}

func (lander *GeoscapeLander) hoverinfo() {
	for _, mr := range lander.mrs {
		if mr.Hovered {
			lander.displayinfo(mr)
			break
		}
	}
}

func (lander *GeoscapeLander) chooselanding() {
	for _, mr := range lander.mrs {
		if mr.Clicked {
			tactical := &TacticalScene{}
			tactical.Region = mr.Region
			tactical.TileSize = lander.TileSize
			tactical.ScreenDims = lander.ScreenDims

			engo.SetScene(tactical, false)
			break
		}
	}
}

func (lander *GeoscapeLander) wipeinfo() {
	if lander.regioninfo != nil {
		derenderentity(lander.world, &lander.regioninfo.BasicEntity)
	}

	lander.regioninfo = nil
}

func (lander *GeoscapeLander) displayinfo(mr *mouseregion) {
	size := lander.TextSize()

	position := func(texture *common.Texture) (float32, float32) {
		return (lander.ScreenWidth - texture.Width()) / 2, lander.ScreenWidth - 10 - size
	}

	msg := fmt.Sprintf("%v (%v,%v)", mr.Region.Class.ShortName(), mr.X, mr.Y)

	hud := hudmsg(msg, size, position)

	lander.regioninfo = hud

	renderentity(lander.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
}

func (lander *GeoscapeLander) Remove(e ecs.BasicEntity) {}

type geotile struct {
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

package colony

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type TacticalScene struct {
	TileSize float32
	ScreenDims

	Region *Region
	world *ecs.World
}

func (scene *TacticalScene) Type() string { return "tactical" }

func (scene *TacticalScene) Preload() { }

func (scene *TacticalScene) Setup(world *ecs.World) {
	common.SetBackground(color.Black)

	scene.world = world
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	hudsys := &TacHudSystem{}
	hudsys.TileSize = scene.TileSize
	hudsys.ScreenDims = scene.ScreenDims
	hudsys.Region = scene.Region

	world.AddSystem(hudsys)

	scene.init()
}

func (scene *TacticalScene) addtile(i, j int) {
	tactile := &TacTile{}

	tactile.BasicEntity = ecs.NewBasic()

	tile := scene.Region.Tiles[strideindex(i, j, scene.Region.Width)]

	tactile.TileComponent = TileComponent{X: i, Y: j, Tile: tile}

	drawable, err := tile.Class.Drawable(scene.TileSize)

	if err != nil {
		panic(err)
	}

	tactile.RenderComponent = rndcomp(drawable)

	fi, fj := float32(i), float32(j)
	x := fi * scene.TileSize
	y := fj * scene.TileSize

	tactile.SpaceComponent = spacecompsz(x, y, scene.TileSize, scene.TileSize)

	mouseentity(scene.world, &tactile.BasicEntity, &tactile.MouseComponent, &tactile.RenderComponent, &tactile.SpaceComponent)
	renderentity(scene.world, &tactile.BasicEntity, &tactile.RenderComponent, &tactile.SpaceComponent)

	for _, system := range scene.world.Systems() {
		switch sys := system.(type) {
		case *TacHudSystem:
			sys.Add(&tactile.BasicEntity, &tactile.MouseComponent, &tactile.TileComponent)
		}
	}
}

func (scene *TacticalScene) init() {
	scene.Region.Class.GenerateTiles()

	for i := 0; i < scene.Region.Width; i++ {
		for j := 0; j < scene.Region.Height; j++ {
			scene.addtile(i, j)
		}
	}

	scene.hudbackground()
}

func (scene *TacticalScene) hudbackground() {
	bg := hudbg(0, scene.ScreenHeight - 50, scene.ScreenWidth, scene.ScreenHeight)

	renderentity(scene.world, &bg.BasicEntity, &bg.RenderComponent, &bg.SpaceComponent)
}


type TacHudSystem struct {
	ScreenDims
	TileSize float32

	Region *Region

	drawn bool

	world *ecs.World

	tiles []mousetile

	tileinfo *HudSection
}

type mousetile struct {
	*common.MouseComponent
	*TileComponent
}

func (hudsys *TacHudSystem) New(w *ecs.World) {
	hudsys.world = w
}

func (hudsys *TacHudSystem) Update(dt float32) {
	hudsys.wipeinfo()

	for _, t := range hudsys.tiles {
		if t.Hovered {
			hudsys.displayinfo(t)

			break
		}
	}
}

func (hudsys *TacHudSystem) Add(b *ecs.BasicEntity, m *common.MouseComponent, t *TileComponent) {
	mt := mousetile{}
	mt.MouseComponent = m
	mt.TileComponent = t

	hudsys.tiles = append(hudsys.tiles, mt)
}

func (hudsys *TacHudSystem) Remove(ecs.BasicEntity) {
}

func (hudsys *TacHudSystem) wipeinfo() {
	if hudsys.tileinfo != nil {
		derenderentity(hudsys.world, &hudsys.tileinfo.BasicEntity)
	}

	hudsys.tileinfo = nil
}

func (hudsys *TacHudSystem) displayinfo(tile mousetile) {
	size := hudsys.TextSize()

	position := func(texture *common.Texture) (float32, float32) {
		return (hudsys.ScreenWidth - texture.Width()) / 2, hudsys.ScreenWidth - 10 - size
	}

	msg := fmt.Sprintf("%v (%v,%v)", tile.Tile.Class.ShortName(), tile.X, tile.Y)

	hud := hudmsg(msg, size, position)

	hudsys.tileinfo = hud

	renderentity(hudsys.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
}

type TacTile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	TileComponent
}

type TileComponent struct {
	X, Y int
	Tile *Tile
}

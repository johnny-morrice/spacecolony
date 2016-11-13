package colony

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type TacticalScene struct {
	TileView

	Region *Region
}

func (ts *TacticalScene) Type() string { return "tactical" }

func (ts *TacticalScene) Preload() { }

func (ts *TacticalScene) Setup(world *ecs.World) {
	common.SetBackground(color.Black)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	tacsys := &TacticalSystem{}
	tacsys.TileView = ts.TileView
	tacsys.Region = ts.Region

	world.AddSystem(tacsys)
}

type TacticalSystem struct {
	TileView

	Region *Region

	drawn bool

	world *ecs.World

	tiles []*TacTile

	tileinfo *HudSection
}

func (tacsys *TacticalSystem) New(w *ecs.World) {
	tacsys.world = w
}

func (tacsys *TacticalSystem) Update(dt float32) {
	if tacsys.drawn {
		tacsys.updatehud()
	} else {
		tacsys.regen()

		tacsys.drawn = true
	}
}

func (tacsys *TacticalSystem) Remove(ecs.BasicEntity) {
}

func (tacsys *TacticalSystem) regen() {
	tacsys.Region.Class.GenerateTiles(&Random{})

	for i := 0; i < tacsys.Region.Width; i++ {
		for j := 0; j < tacsys.Region.Height; j++ {
			tacsys.addtile(i, j)
		}
	}
}

func (tacsys *TacticalSystem) addtile(i, j int) {
	tactile := &TacTile{}

	tactile.BasicEntity = ecs.NewBasic()

	tile := tacsys.Region.Tiles[strideindex(i, j, tacsys.Region.Width)]

	tactile.TileComponent = TileComponent{X: i, Y: j, Tile: tile}

	drawable, err := tile.Class.Drawable(tacsys.TileSize)

	if err != nil {
		panic(err)
	}

	tactile.RenderComponent = common.RenderComponent{
		Drawable: drawable,
		Scale: engo.Point{X: 1, Y: 1},
	}

	fi, fj := float32(i), float32(j)
	x := (fi * tacsys.TileSize) + tacsys.OffsetX
	y := (fj * tacsys.TileSize) + tacsys.OffsetY

	tactile.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: x, Y: y},
		Width: tacsys.TileSize,
		Height: tacsys.TileSize,
	}

	tacsys.tiles = append(tacsys.tiles, tactile)

	mouseentity(tacsys.world, &tactile.BasicEntity, &tactile.MouseComponent, &tactile.RenderComponent, &tactile.SpaceComponent)
	renderentity(tacsys.world, &tactile.BasicEntity, &tactile.RenderComponent, &tactile.SpaceComponent)
}

func (tacsys *TacticalSystem) updatehud() {
	tacsys.wipeinfo()

	for _, tactile := range tacsys.tiles {
		if tactile.Hovered {
			tacsys.displayinfo(tactile)

			break
		}
	}
}

func (tacsys *TacticalSystem) wipeinfo() {
	if tacsys.tileinfo != nil {
		derenderentity(tacsys.world, &tacsys.tileinfo.BasicEntity)
	}

	tacsys.tileinfo = nil
}

func (tacsys *TacticalSystem) displayinfo(tactile *TacTile) {
	size := (tacsys.ScreenHeight - tacsys.ViewSquareSize) / 12

	position := func(texture *common.Texture) (float32, float32) {
		return (tacsys.ScreenWidth - texture.Width()) / 2, tacsys.ScreenWidth - 10 - size
	}

	msg := fmt.Sprintf("%v (%v,%v)", tactile.Tile.Class.ShortName(), tactile.X, tactile.Y)

	hud := hudmsg(msg, size, position)

	tacsys.tileinfo = hud

	renderentity(tacsys.world, &hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
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

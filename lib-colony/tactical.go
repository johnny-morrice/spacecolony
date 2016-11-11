package colony

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type TacticalScene struct {
	ScreenDims

	Region *Region
}

func (ts *TacticalScene) Type() string { return "tactical" }

func (ts *TacticalScene) Preload() { }

func (ts *TacticalScene) Setup(world *ecs.World) {
	common.SetBackground(color.Black)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	tacsys := &TacticalSystem{}
	tacsys.TileView = NewTileView(ts.ScreenDims)

	world.AddSystem(tacsys)
}

type TacticalSystem struct {
	TileView

	Region *Region

	drawn bool

	world *ecs.World
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
	tacsys.Region.Init(&Random{})
}

func (tacsys *TacticalSystem) updatehud() {

}

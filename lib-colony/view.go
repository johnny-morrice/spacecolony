package colony

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type TileView struct {
	TileSize float32
	ViewSquareSize float32
	OffsetX float32
	OffsetY float32
	ScreenWidth float32
	ScreenHeight float32
}

func NewTileView(sd ScreenDims) TileView {
	tv := TileView{}

	tv.ScreenWidth = float32(sd.Width)
	tv.ScreenHeight = float32(sd.Height)

	bound := tv.ScreenWidth
	bigger := tv.ScreenHeight
	if bound > bigger {
		bigger, bound = bound, bigger
	}

	margin := bound / 4
	tv.ViewSquareSize = bound - margin
	tv.OffsetX = (tv.ScreenWidth - tv.ViewSquareSize) / 2
	tv.OffsetY = (tv.ScreenHeight - tv.ViewSquareSize) / 2

	return tv
}

type HudSection struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func hudmsg(msg string, size float32, position func(*common.Texture) (float32, float32)) *HudSection {
	texture, err := basictext(msg, size)

	if err != nil {
		panic(err)
	}

	x, y := position(texture)

	hud := &HudSection{}
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

	return hud
}

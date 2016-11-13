package colony

import (
	"engo.io/engo"
)

type GameOptions struct {
        EngineOptions
}

type ScreenDims struct {
	Width uint
        Height uint
}

type EngineOptions struct {
	ScreenDims
        FPS uint
        Samples uint
        Vsync bool
        Fullscreen bool
}

func Play(opts GameOptions) {
	runopts := engo.RunOptions{
		Title: "Space Colony",
		Width:  int(opts.Width),
		Height: int(opts.Height),
                FPSLimit: int(opts.FPS),
                MSAA: int(opts.Samples),
                VSync: opts.Vsync,
                Fullscreen: opts.Fullscreen,
	}

	scene := &GeoscapeScene{}
	scene.TileView = NewTileView(opts.ScreenDims)

	engo.Run(runopts, scene)
}

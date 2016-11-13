package colony

import (
	"engo.io/engo"
)

type GameOptions struct {
        EngineOptions
}

type ScreenOpts struct {
	Width uint
        Height uint
}

type EngineOptions struct {
	ScreenOpts
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
	scene.ScreenDims = ScreenDims{
		ScreenWidth: float32(opts.Width),
		ScreenHeight: float32(opts.Height),
	}

	engo.Run(runopts, scene)
}

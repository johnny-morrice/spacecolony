package colony

import (
	"engo.io/engo"
)

type GameOptions struct {
        EngineOptions
}

type EngineOptions struct {
        Width uint
        Height uint
        FPS uint
        Samples uint
        Vsync bool
        Fullscreen bool
}

func Play(gopts GameOptions) {
	eopts := engo.RunOptions{
		Title: "Space Colony",
		Width:  int(gopts.Width),
		Height: int(gopts.Height),
                FPSLimit: int(gopts.FPS),
                MSAA: int(gopts.Samples),
                VSync: gopts.Vsync,
                Fullscreen: gopts.Fullscreen,
	}

	scene := &geoscapeScene{}
	scene.EngineOptions = gopts.EngineOptions

	engo.Run(eopts, scene)
}

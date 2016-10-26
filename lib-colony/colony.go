package colony

import (
	"engo.io/engo"
	"engo.io/ecs"
)

type colonyScene struct {}

// Type uniquely defines your game type
func (*colonyScene) Type() string { return "colony" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*colonyScene) Preload() {}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*colonyScene) Setup(*ecs.World) {}

type GameOptions struct {
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

	engo.Run(eopts, &colonyScene{})
}

package colony

import (
        "image/color"
	"path/filepath"
	"os"

	"engo.io/engo"
        "engo.io/engo/common"
	"engo.io/ecs"
)

type colonyScene struct {
	DisplayOptions
}

func (cs *colonyScene) Type() string { return "colony" }

func (cs *colonyScene) Preload() {
	var matches []string

	err := filepath.Walk("assets/png", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		_, file := filepath.Split(path)

		matches = append(matches, filepath.Join("png", file))

		return nil
	})

	if err != nil {
		panic(err)
	}

	err = engo.Files.Load(matches...)

	if err != nil {
		panic(err)
	}
}

func (cs *colonyScene) Setup(world *ecs.World) {
        common.SetBackground(color.Black)

        world.AddSystem(&common.RenderSystem{})
        world.AddSystem(&common.MouseSystem{})

	geoscape := &GeoscapeSystem{}
	geoscape.Tilesize = float32(cs.DisplayOptions.Tilesize)

        world.AddSystem(geoscape)
}
type GameOptions struct {
        EngineOptions
	DisplayOptions
}

type DisplayOptions struct {
	Tilesize uint
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

	scene := &colonyScene{}
	scene.DisplayOptions = gopts.DisplayOptions

	engo.Run(eopts, scene)
}

package colony

import (
        "image/color"
	"path/filepath"
	"os"

	"engo.io/engo"
        "engo.io/engo/common"
	"engo.io/ecs"
)

type geoscapeScene struct {
	DisplayOptions
	EngineOptions
}

func (gs *geoscapeScene) Type() string { return "geoscape" }

func (gs *geoscapeScene) Preload() {
	err := loadAllAssets()

	if err != nil {
		panic(err)
	}
}

func (gs *geoscapeScene) Setup(world *ecs.World) {
        common.SetBackground(color.Black)

        world.AddSystem(&common.RenderSystem{})
        world.AddSystem(&common.MouseSystem{})

	geoscape := &GeoscapeSystem{}
	geoscape.Tilesize = float32(gs.Tilesize)
	geoscape.ScreenWidth = float32(gs.Width)
	geoscape.ScreenHeight = float32(gs.Height)

        world.AddSystem(geoscape)
}

func loadAllAssets() error {
	var matches []string

	err := filepath.Walk("assets/sprite", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		_, file := filepath.Split(path)

		matches = append(matches, filepath.Join("sprite", file))

		return nil
	})

	if err != nil {
		return err
	}

	return engo.Files.Load(matches...)
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

	scene := &geoscapeScene{}
	scene.DisplayOptions = gopts.DisplayOptions
	scene.EngineOptions = gopts.EngineOptions

	engo.Run(eopts, scene)
}

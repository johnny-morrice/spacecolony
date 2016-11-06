package colony

import (
	"image/color"
	"path/filepath"
	"os"

	"engo.io/engo"
	"engo.io/engo/common"
)

func loadAllAssets() error {
	assets := []string{
		"sprite",
		"font",
	}

	for _, dir := range assets {
		err := loadAssetDir(dir)

		if err != nil {
			return err
		}
	}

	return nil
}

func loadAssetDir(dir string) error {
	var matches []string

	err := filepath.Walk(filepath.Join("assets", dir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		_, file := filepath.Split(path)

		matches = append(matches, filepath.Join(dir, file))

		return nil
	})

	if err != nil {
		return err
	}

	return engo.Files.Load(matches...)
}

func stdfont() *common.Font {
	fnt := &common.Font{}
	fnt.URL = "font/monospace.ttf"
	fnt.BG = color.Black
	fnt.FG = color.White

	return fnt
}

func basicText(text string, size float64) (*common.Texture, error) {
	fnt := stdfont()
	fnt.Size = size

	err := fnt.CreatePreloaded()

	if err != nil {
		return nil, err
	}

	texture := fnt.Render(text)

	return &texture, nil
}

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

func monospace(size float64) (*common.Font, error) {
	fnt := &common.Font{}
	fnt.URL = "font/monospace.ttf"
	fnt.Size = size
	fnt.BG = color.Black
	fnt.FG = color.White

	err := fnt.CreatePreloaded()

	if err != nil {
		return nil, err
	}

	return fnt, nil
}

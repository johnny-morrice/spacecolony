package colony

import (
	"image"
	"image/color"
	"image/draw"

	"engo.io/engo/common"
)

func uniformimg(c color.NRGBA, width, height float32) *image.NRGBA {
	iw, ih := int(width), int(height)

	bounds := image.Rect(0, 0, iw, ih)

	source := image.NewUniform(c)
	out := image.NewNRGBA(bounds)
	draw.Draw(out, bounds, source, image.ZP, draw.Src)

	return out
}

func imgtexture(img *image.NRGBA) *common.Texture {
	obj := common.NewImageObject(img)

	texture := common.NewTextureSingle(obj)
	return &texture
}

func basictext(text string, size float32) (*common.Texture, error) {
	fnt := stdfont()
	fnt.Size = float64(size)

	err := fnt.CreatePreloaded()

	if err != nil {
		return nil, err
	}

	texture := fnt.Render(text + " ")

	return &texture, nil
}

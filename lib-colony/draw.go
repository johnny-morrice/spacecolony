package colony

import (
	"image"
	"image/color"
	"image/draw"

	"engo.io/engo/common"
)

func uniformimg(c color.NRGBA, width, height float64) *image.NRGBA {
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

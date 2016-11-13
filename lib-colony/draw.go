package colony

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"engo.io/ecs"
	"engo.io/engo/common"

	"github.com/ojrac/opensimplex-go"
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

func noiseimg(seed int64, size int) *image.NRGBA {
	rect := image.Rectangle{Max: image.Point{X:size, Y: size}}
	img := image.NewNRGBA(rect)

	noise := opensimplex.NewWithSeed(seed)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			shade := noise.Eval2(float64(i), float64(j))
			gr := uint8(math.Floor(shade + 1) * 125)
			col := color.NRGBA{R: gr, G: gr, B: gr, A: 255}

			img.SetNRGBA(i, j, col)
		}
	}

	return img
}

type ScreenDims struct {
	ScreenWidth float32
	ScreenHeight float32
}

func (sd ScreenDims) TextSize() float32 {
	return sd.ScreenWidth / 40
}

type CenterTiles struct {
	ViewSquareSize float32
	VSMinX float32
	VSMinY float32
	VSMaxX float32
	VSMaxY float32
}

func NewCenterTiles(sd ScreenDims) CenterTiles {
	tv := CenterTiles{}

	bound := sd.ScreenWidth
	bigger := sd.ScreenHeight
	if bound > bigger {
		bigger, bound = bound, bigger
	}

	margin := bound / 4
	tv.ViewSquareSize = bound - margin
	tv.VSMinX = (sd.ScreenWidth - tv.ViewSquareSize) / 2
	tv.VSMinY = (sd.ScreenHeight - tv.ViewSquareSize) / 2

	return tv
}

type HudSection struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func hudmsg(msg string, size float32, position func(*common.Texture) (float32, float32)) *HudSection {
	texture, err := basictext(msg, size)

	if err != nil {
		panic(err)
	}

	x, y := position(texture)

	hud := &HudSection{}

	hud.BasicEntity = ecs.NewBasic()

	hud.SpaceComponent = spacecompsz(x, y, texture.Width(), texture.Height())

	hud.RenderComponent = rndcomp(texture)

	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(2)

	return hud
}

func hudbg(xmin, ymin, xmax, ymax float32) *HudSection {
	black := color.NRGBA{A: 255}

	img := uniformimg(black, xmax - xmin, ymax - ymin)

	texture := imgtexture(img)

	hud := &HudSection{}

	hud.BasicEntity = ecs.NewBasic()

	hud.SpaceComponent = spacecomprect(xmin, ymin, xmax, ymax)

	hud.RenderComponent = rndcomp(texture)

	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

	return hud
}

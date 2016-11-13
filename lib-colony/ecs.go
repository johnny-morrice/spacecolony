package colony

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

func renderentity(w *ecs.World, b *ecs.BasicEntity, r *common.RenderComponent, s *common.SpaceComponent) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(b, r, s)
		}
	}
}

func derenderentity(w *ecs.World, b *ecs.BasicEntity) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(*b)
		}
	}
}

func mouseentity(w *ecs.World, b *ecs.BasicEntity, m *common.MouseComponent, r *common.RenderComponent, s *common.SpaceComponent) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(b, m, s, r)
		}
	}
}

func rndcomp(drawable common.Drawable) common.RenderComponent {
	return common.RenderComponent{
		Drawable: drawable,
		Scale: engo.Point{X: 1, Y: 1},
	}
}

func spacecomprect(xmin, ymin, xmax, ymax float32) common.SpaceComponent {
	return common.SpaceComponent{
		Position: engo.Point{X: xmin, Y: ymin},
		Width: xmin + xmax,
		Height: ymin + ymax,
	}
}

func spacecompsz(x, y, width, height float32) common.SpaceComponent {
	return common.SpaceComponent{
		Position: engo.Point{X: x, Y: y},
		Width: width,
		Height: height,
	}
}

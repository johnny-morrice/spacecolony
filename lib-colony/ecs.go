package colony

import (
	"engo.io/engo/common"
	"engo.io/ecs"
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

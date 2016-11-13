package colony

import (
	"testing"
)

func TestPlanetInit(t *testing.T) {
	p := &Planet{}
	p.Width = 100
	p.Height = 100

	p.Init(&math.Rand{})

	if len(p.Tiles) != p.Width * p.Height {
		t.Error("Unexpected number of tiles")
	}
}

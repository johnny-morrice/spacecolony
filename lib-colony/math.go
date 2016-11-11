package colony

import (
	"math"
)

type Bounds struct {
        XMin int
        XMax int
        YMin int
        YMax int
}

func evenfloor(x float32) float32 {
	x = float32(math.Floor(float64(x)))

	if int(x) % 2 == 0 {
		return x
	} else {
		return x - 1
	}
}

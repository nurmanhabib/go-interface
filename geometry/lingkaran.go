package geometry

import (
    "fmt"
    "math"
)

type Lingkaran struct {
	Radius float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.Radius * l.Radius
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.Radius
}

func (l Lingkaran) Print() {
    fmt.Println(l.Luas())
}

package geometry

import "fmt"

type Persegi struct {
	Lebar int
	Panjang int
}

func (p Persegi) Luas() float64 {
	return float64(p.Lebar * p.Panjang)
}

func (p Persegi) Keliling() float64 {
	return float64(2 * p.Lebar + p.Panjang)
}

func (p Persegi) Print() {
    fmt.Println(p.Luas())
}

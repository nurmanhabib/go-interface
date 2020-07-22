package geometry

type Balok struct {
    Persegi
    Tinggi float64
}

func (b Balok) Volume() float64 {
    return float64(b.Panjang) * float64(b.Lebar) * b.Tinggi
}

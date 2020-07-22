package contract

type GeometryInterface interface {
    Printable
    Luas() float64
    Keliling() float64
    Print()
}

type TigaDimensi interface {
    Volume() float64
}

type Printable interface {
    Print()
}

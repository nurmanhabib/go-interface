package main

import (
	"github.com/nurmanhabib/go-interface/geometry"
)

func main() {
	lingkaran := geometry.Lingkaran{Radius: 5}
	persegi := geometry.Persegi{Panjang: 6, Lebar: 4}

	geometry.CetakLuasDanKeliling(lingkaran)
	geometry.CetakLuasDanKeliling(persegi)
}

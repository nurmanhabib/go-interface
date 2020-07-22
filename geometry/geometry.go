package geometry

import (
    "fmt"
    "github.com/nurmanhabib/go-interface/contract"
)

func CetakLuasDanKeliling(g contract.GeometryInterface) {
    fmt.Printf("Luas = %f", g.Luas())
    fmt.Println()
    fmt.Printf("Keliling = %f", g.Keliling())
    fmt.Println()
}

package main
import "fmt"
func main() {
    var hora, min, seg int
    fmt.Scan(&hora, &min, &seg)

    seg = seg + 1 

    if seg == 60 {
        seg = 00
        min = min + 1
    }
    if min == 60 {
        min = 00 
        hora = hora + 1
    }
    if hora == 24 {
        hora = 00
        min = 00
        seg = 00
    }
    fmt.Printf("%02d %02d %02d\n", hora, min, seg)
}

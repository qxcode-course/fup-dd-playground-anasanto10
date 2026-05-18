package main
import "fmt"

func existe(arr[]int, valor int) bool {
    for _, elem := range arr {
        if elem ==  valor {
            return true
        }
    }
    return false
}

func estaContido (v1, v2 []int) bool {
    for _, elem := range v1 {
        if !existe(v2, elem) {
            return false
        }
    }
    return true
}
    
    func main() {
    var N int
    fmt.Scan(&N)
    v1 := make([]int, N)
    for i := range v1{
        fmt.Scan(&v1[i])
    }

    var M int
    fmt.Scan(&M)
    v2 := make([]int, M)
    for i := range v2 {
        fmt.Scan(&v2[i])
    }

    if estaContido(v1, v2) {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }

}

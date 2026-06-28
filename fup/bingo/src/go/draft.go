package main
import "fmt"
func main() {
    numeros := make([]int, 6)
    for i := range numeros {
        fmt.Scan(&numeros[i])
    }

    matriz := [4][4]int {
        {1, 9, 27, 23},
        {34, 20, 37, 47},
        {30, 87, 55, 69},
        {13, 60, 99, 66},
    }
    acertos := 0

    for _, num := range numeros {
        for i := 0; i < 4 ; i++ {
            for j := 0; j < 4; j++ {
                if num == matriz[i][j] {
                    acertos++
                }
            }
        }
    }
    fmt.Println(acertos)
}
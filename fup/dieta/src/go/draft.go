package main
import "fmt"
func main() {
    var dias int
    fmt.Scan(&dias)
    calorias := make([]int, dias)

    soma := 0

    for i := 0; i < dias; i++ {
        fmt.Scan(&calorias[i])

        soma += calorias[i]
    }

    media := float64(soma) / float64(dias)
    fmt.Printf("%.1f\n", media)
}



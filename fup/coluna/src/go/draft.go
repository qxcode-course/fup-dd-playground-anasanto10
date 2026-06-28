package main
import "fmt"
func main() {
    var n int
    fmt.Scan(n)

    matriz := make([][]int, n)
    for i := 0; i < n; i++ {
        matriz[i] = make([]int, n)

        for j := 0; j < n; j++{
            fmt.Scan(&matriz[i][j])
        }
    }

    maiorVal := -1
    indiceMai := 0

    for coluna := 0; coluna < n; coluna++ {
        soma := 0

        for linha := 0; linha < n; linha++{
            soma += matriz[linha][coluna] * matriz[linha][coluna]
        }

        if soma > maiorVal {
            maiorVal = soma
            indiceMai = coluna
        }
    }
    fmt.Println(indiceMai)
}
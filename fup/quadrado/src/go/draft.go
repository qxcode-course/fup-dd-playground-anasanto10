package main
import "fmt"
func main() {
    var matriz [3][3]int

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&matriz[i][j])
        }
    }

    soma := matriz[0][0] + matriz[0][1] + matriz[0][2]
    magico := true

    for i := 0; i < 3; i++ {
        somaLinha := 0
        for j := 0; j < 3; j++ {
            somaLinha += matriz[i][j]
        }
        if somaLinha != soma {
            magico = false
        }
    }

    for j := 0; j < 3; j++ {
        somaColuna := 0
        for i := 0; i < 3; i++ {
            somaColuna += matriz[i][j]
        }
        if somaColuna != soma {
            magico = false
        }
    }

    somaDiagonal := 0 
    for i := 0; i < 3; i++ {
        somaDiagonal += matriz[i][2-1]
    }
    if somaDiagonal != soma {
        magico = false
    }

    if magico {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
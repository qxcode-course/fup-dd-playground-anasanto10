package main
import "fmt"
func main() {
    var nota1, nota2, nota_final float32
    var media float32
    fmt.Scan(&nota1, &nota2, &nota_final)
    
    media = (nota1 + nota2) / 2

    if (media >= 7){
        fmt.Println("aprovado")
    } else if (media >= 4 && media <= 7){
        mf := (media + nota_final) / 2
        if (mf >= 5){
            fmt.Println("aprovado na final")
        } else if (mf <5){
            fmt.Println("reprovado na final")
        }
    } else if (media < 4){
        fmt.Println("reprovado")
    }
}

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
        
    }
}

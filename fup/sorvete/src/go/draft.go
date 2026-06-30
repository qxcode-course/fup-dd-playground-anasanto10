package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    frase := scanner.Text()
    
    resposta := 0 
    for letra := byte('a'); letra <= 'z'; letra++ {

        comum := true
        inicio := 0

        for inicio < len(frase) {

            achou := false
            fim := inicio

            for fim < len(frase) && frase[fim] != ' ' {
                if frase[fim] == letra {
                    achou = true
                }
                fim++
            }

            if !achou {
                comum = false
                break
            }
            inicio = fim + 1
        }
        if comum {
            resposta++
        }
    }

    fmt.Println(resposta)
}
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
    scanner.Scan()

    palavra := ""

    for i := 0; i < len(frase); i++ {
        if frase[i] != ' ' {
            palavra += string(frase[i])          
        } else {
            fmt.Print(palavra, " ", palavra, " ")
            palavra = ""
        }
    }

    if palavra != "" {
        fmt.Print(palavra, " ", palavra)
    }
    
    fmt.Println()
}
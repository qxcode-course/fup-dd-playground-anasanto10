package main
import (
    "bufio"
    "os"
    "fmt"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    scanner.Scan()

    var vogais []rune
    var consoantes []rune

    for _, letras := range line {

        if letras != ' ' {
        if letras == 'a' || letras == 'e' || letras == 'i' || letras == 'o' || letras == 'u' {
            vogais = append(vogais, letras)
        } else {
            consoantes = append(consoantes, letras)
        }
    }
}
    fmt.Println(string(vogais))
    fmt.Println(string(consoantes))
}

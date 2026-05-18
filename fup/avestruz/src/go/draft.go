package main
import (
    "fmt"
    "bufio"
    "os"
)

func toLower(value rune) rune {
    if value >= 'A' && value <= 'Z' {
        return value + ('a' - 'A')
    }
    return value
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    scanner.Scan()
    palavra := scanner.Text()
    letra := rune(palavra[0])

    cont := 0
    for _, c := range line {
        if toLower(c) == toLower(letra) {
            cont += 1
        }
    }
    fmt.Println(cont)
}

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
    var letra rune
    fmt.Sscanf(scanner.Text(), "%c", &letra)

    cont := 0
    for _, c := range line {
        if c == letra {
            cont += 1
        }
    }
    fmt.Println(cont)
}

package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    scanner.Scan()

    trecho := scanner.Text()
    scanner.Scan()

    cont := strings.Count(line, trecho)
    
    fmt.Println(cont)
}
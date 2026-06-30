package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    scanner.Scan()

    pal := scanner.Text()
    scanner.Scan()

    subs := scanner.Text()
    scanner.Scan()

    i := 0

    for i <= len(line)-len(pal){
        if line[i:i+len(pal)] == pal {
            fmt.Print(subs)
            i += len(pal)
        } else {
            fmt.Printf("%c", line[i])
            i++
        }
    }

    for i < len(line) {
        fmt.Printf("%c", line[i])
        i++
    }
    fmt.Println()
}

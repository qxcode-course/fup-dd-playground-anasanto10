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

    for i := 0; i <= len(line)-len(pal); i++ {
        
    }
}

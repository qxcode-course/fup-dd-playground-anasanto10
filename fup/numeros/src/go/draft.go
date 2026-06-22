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

    soma := 0
    for i := 0; i < len(line); i++ {
        if line[i] >= '0' && line[i] <= '9' {
            num := 0

            if i < len(line) && line[i] >= '0' && line[i] <= '9' {
                num = num*10 + int(line[i]-'0')
                i++
            }
            soma += num
        }
    }
    fmt.Println(soma)
}
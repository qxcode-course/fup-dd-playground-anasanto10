package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func Substring(s string, ind, qtd int) string {
    if ind < 0 || ind >= len(s) || qtd <=0 {
        return ""
    }
    fim := ind + qtd

    if fim > len(s) {
        fim = len(s)
    }

    return s[ind:fim]
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    s := scanner.Text()

    scanner.Scan()
    ind, _ := strconv.Atoi(scanner.Text())

    scanner.Scan()
    qtd, _ := strconv.Atoi(scanner.Text())

    fmt.Print(Substring(s, ind, qtd))
    fmt.Println()
}
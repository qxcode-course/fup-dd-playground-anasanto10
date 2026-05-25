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

    l := []rune(line)

    for i:= 0; i < len(line); i++ {
    
        if  l[i] == 'A' || l[i] == 'a' ||
            l[i] == 'E' || l[i] == 'e'||
            l[i] == 'I' || l[i] == 'i' ||
            l[i] == 'O' || l[i] == 'o' ||
            l[i] == 'U' || l[i] == 'u' {
           l[i] = 'v'
        } else if l[i] != ' '{
           l[i] = 'c'
        }
    }
    for i := 0; i < len(line); i++ {
        fmt.Printf("%c", l[i])
    }

   fmt.Println()
}

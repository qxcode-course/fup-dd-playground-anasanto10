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

   for i := 0; i < len(frase); i++ {
       if frase[i] == '#' {
        fmt.Println()
       } else if frase[i] == ';' {
        if i != len(frase)-1 {
            fmt.Println()
        }
       } else {
        fmt.Printf("%c", frase[i])
       }
   }
   fmt.Println()
}
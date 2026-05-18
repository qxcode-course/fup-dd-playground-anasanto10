package main
import "fmt"
func main() {
    var palavra string
    var delta int
    fmt.Scan(&palavra, &delta)
    var letra int = int(palavra[0])

    letra -= 'a' //0 e 25
    letra += delta

    if letra < 0{
        letra += 26
    } else if letra >= 26 {
        letra -= 26
    }

    letra += 'a'

    fmt.Printf("%c\n", letra)
}

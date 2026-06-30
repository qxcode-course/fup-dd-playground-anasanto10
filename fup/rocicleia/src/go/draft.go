package main
import "fmt"
func main() {
    var p1, p2 string
    fmt.Scan(&p1, &p2)

    cont := make([]int, 26)

    for i := 0; i < len(p1); i++ {
        cont[p1[i]-'a']++
    }

    for i := 0; i < len(p2); i++ {
        cont[p2[i]-'a']--
    }

    anagrama := true

    for i := 0; i < 26; i++{
        if cont[i] != 0 {
            anagrama = false
            break
        }
    }

    if anagrama {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}
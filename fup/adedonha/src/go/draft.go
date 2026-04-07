package main

import "fmt"

func main() {
    var i1, i2, i3 int
    fmt.Scan(&i1, &i2, &i3)

    soma := i1 + i2 + i3

    if soma == 0 {
        fmt.Println("joguem de novo")
    } else {
        letra := 'a' + rune((soma-1)%26)
        fmt.Printf("%c\n", letra)
    }
}
package main
import "fmt"

func criarEpreencher() []int {
    qtd := 0
    fmt.Scan(&qtd)
    arr := make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    return arr
}

func main() {
    arr := criarEpreencher()
    animal := []int{}

    for i := 0; i < len(arr); i++ {
        existe := false

        for j := 0; j < len(animal); j++ {
            if arr[i] == animal[j] {
                existe = true
            } 
        }
        if !existe {
                animal = append(animal, arr[i])
            }
    }
    for i := 0; i < len(animal)-1; i++ {
        for j := 0; j < len(animal)-1; j++ {
            if animal[j] > animal[j+1] {
                animal[j], animal[j+1] = animal[j+1], animal[j]
            }
        }
    }
    for i := 0; i < len(animal); i++ {
        if i > 0{
            fmt.Print(" ")
        }
        fmt.Print(animal[i])
    }
    fmt.Println()
}
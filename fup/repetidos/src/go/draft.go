package main
import "fmt"
func main() {
    var p, n int 
    fmt.Scan(&p, &n)
    vet := make([]int, n)

    for i := 0; i < n; i++{
        fmt.Scanf("%d", &vet[i])
    }

    cont := 0

    for i := 0; i < n; i++{
        if vet[i] == p {
            cont ++
        }
    }
    
    fmt.Println(cont)
}

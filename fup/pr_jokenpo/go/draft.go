package main
import "fmt"

// 0 pedra
// 1 papel
// 2 tesoura
func main() {
    const pedra int = 0
    const papel int = 1
    const tesoura int = 2
    
    var jog1, jog2 int
    var vit1, vit2 int

    for {
        if vit1 > 2 || vit2 > 2{
            break
        }
    fmt.Println("Jog1: digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog1)
    fmt.Println("Jog2:digite 0(pedra), 1(papel), 2(tesoura):")
    fmt.Scan(&jog2)

    if jog1 == jog2 [
        fmt.Println("empate")
    ] else if (jog1 == pedra && jog2 == tesoura) ||
              (jog1 == papel && jog2 == pedra) ||
              (jog1 == tesoura && jog2 == papel){
                fmt.Println("jog1 ganhou")
                vit1 += 1
              } else {
                fmt.Println("jog2 ganhou")
                vit2 += 1
              }
            }

            if vit1 > vit2 {
                fmt.Println("jog1 ganhou a partida")
            } else {
                fmt.Println("jog2 ganhou a partida")
            }

}

package main
import "fmt"
func main() {
    var jog1, jog2 string
    fmt.Scan(&jog1, &jog2)

    if (jog1 == "R" && jog2 == "S" || jog1 == "P" && jog2 == "R" || jog1 == "S" && jog2 == "P"){
        fmt.Println("jog1")
    } else if (jog2 == "R" && jog1 =="S" || jog2 == "P" && jog1 == "R" ||  jog2 == "S" && jog1 == "P"){
        fmt.Println("jog2")
    } else if(jog1 == jog2){
        fmt.Println("empate")
    }
}

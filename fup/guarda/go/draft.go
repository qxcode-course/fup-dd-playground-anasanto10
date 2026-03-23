package main
import "fmt"
func main() {
    var wifi, login, admin int16
    fmt.Scan (&wifi, &login, &admin)
    wifiBool := wifi == 1
    loginBool := login == 1
    adminBool := admin == 1
    if ( !wifiBool ){
     fmt.Println("you must connect to wifi")
    } else if ( !loginBool ){
        fmt.Println("you need to login first")
    } else if ( !adminBool ){
        fmt.Println("you must to login as admin")
    } else {
        fmt.Println("done")
    }
    
}

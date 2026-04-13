package main
import "fmt"
func main() {
	var cap int
	fmt.Scan(&cap)
	var onibus, mov int

	for {
		qtd, _ := fmt.Scan(&mov)
		if qtd == 0 {
			break
		}
	      onibus += mov

		if onibus == 0 {
			fmt.Println("vazio")
		} else if onibus >= cap*2 {
			fmt.Println("hora de partir")
			break
		} else if onibus >= cap {
			fmt.Println("lotado")
		} else if onibus < cap {
			fmt.Println("ainda cabe")
		}
	}
}
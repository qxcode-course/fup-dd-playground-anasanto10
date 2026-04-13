package main
import "fmt"
func palindromo(n int) int {
	original := n
	invertido := 0

	for n > 0 {
		digito := n % 10
		invertido = invertido*10 + digito
		n = n / 10
	}

	if original == invertido {
		return 1
	}
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(palindromo(n))
}
package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	scanner.Scan()

	r := []rune(line)

	i := 0
	j := len(r) - 1

	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	line = string(r)
	fmt.Println(line)
}

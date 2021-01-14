package main
import (
	"fmt"
	"bufio"
	"os"
)
//hw
func min2(a, b int) int {
	if (a > b) {
		return b
	}
	return a
}
func min(a, b, c, d int) int {
	return min2(a, min2(b, min2(c, d)))
}
func main() {
	cin := bufio.NewReader(os.Stdin)
	var a, b, c, d int
	fmt.Fscan(cin, &a, &b, &c, &d)
	fmt.Print(min(a, b, c, d))
}

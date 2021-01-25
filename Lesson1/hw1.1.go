package main

import (
	"bufio"
	"fmt"
	"os"
)

//hw
func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b, c, d int) int {
	return min2(a, min2(b, min2(c, d)))
}
func max(a, b, c, d int) int {
	return max2(a, max2(b, max2(c, d)))
}
func main() {
	cin := bufio.NewReader(os.Stdin)
	var a, b, c, d int
	fmt.Fscan(cin, &a, &b, &c, &d)
	fmt.Println("Max = ", max(a, b, c, d))
	fmt.Println("Min = ", min(a, b, c, d))
}

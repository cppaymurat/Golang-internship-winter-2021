package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	var a int64
	fmt.Fscan(cin, &a)
	a /= 10
	if a == 0 {
		fmt.Print(25)
	} else {
		fmt.Print(a*(a+1), "25")
	}
}

package main
import (
	"fmt"
	"bufio"
	"os"
)

const N = 100000

var dp [N]int

func fib(n int) int {
	if n == 1 {
		dp[1] = 1
		return dp[1]
	}
	if n == 2 {
		dp[2] = 1
		return dp[2]
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = fib(n - 1) + fib(n - 2)
	return dp[n]
}
func main() {
	cin := bufio.NewReader(os.Stdin)
	var x int
	fmt.Fscan(cin, &x)
	fmt.Print(fib(x))
}

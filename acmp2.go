package main
import (
    "fmt"
    "bufio"
    "os"
)
 
func main() {
    cin := bufio.NewReader(os.Stdin)
    var a int
    fmt.Fscan(cin, &a)
    res := 0
    if a < 1 {
        for i := a; i <= 1; i++ {
            res += i
        }
    } else {
        for i := 1; i <= a; i++ {
            res += i
        }
    }
    fmt.Print(res)
}
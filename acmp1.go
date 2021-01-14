apackage main
import (
    "fmt"
    "bufio"
    "os"
)
 
func main() {
    cin := bufio.NewReader(os.Stdin)
    var a, b int
    fmt.Fscan(cin, &a, &b)
    fmt.Print(a + b)
}
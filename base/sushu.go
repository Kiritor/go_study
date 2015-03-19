package main
import (
    "fmt"
    "math"
)
func main() {
    var i, j, n int
    var a [100001]int
    for i = 1; i <= 100000; i++ {
        a[i] = i
    }
    a[1] = 0
    for i = 2; i < int(math.Sqrt(100000)); i++ {
        for j = i + 1; j <= 100000; j++ {
            if (a[i] != 0) && (a[j] != 0) {
                if a[j]%a[i] == 0 {
                    a[j] = 0
                }
            }
        }
    }
    fmt.Println()
    for i, n = 1, 0; i <= 100000; i++ {
        if a[i] != 0 {
            fmt.Print(a[i], "\t")
            n++
        }
        if n == 10 {
            fmt.Println()
            n = 0
        }
    }
}

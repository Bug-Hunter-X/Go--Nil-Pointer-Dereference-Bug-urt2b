```go
package main

import "fmt"

func main() {
    var x int
    fmt.Println(x) // Output: 0
    x = 10
    fmt.Println(x) // Output: 10
    x = x + 10
    fmt.Println(x) // Output: 20
    x += 10
    fmt.Println(x) // Output: 30
    fmt.Println(&x) // Output: 0xc0000140a8

    var y *int
    y = &x
    fmt.Println(y)
    fmt.Println(*y) // Output: 30
    *y = 40
    fmt.Println(*y) // Output: 40
    fmt.Println(x) // Output: 40
    y = nil
    // fmt.Println(*y) // This will cause a runtime error: panic: runtime error: invalid memory address or nil pointer dereference
}
```
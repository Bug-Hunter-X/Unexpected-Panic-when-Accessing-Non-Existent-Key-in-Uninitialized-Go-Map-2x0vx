```go
package main

import "fmt"

func main() {

    var m map[string]int
    fmt.Println(m["a"]) // this will not print 0, but will panic
}
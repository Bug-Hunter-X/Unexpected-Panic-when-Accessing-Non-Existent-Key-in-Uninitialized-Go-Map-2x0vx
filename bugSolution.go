```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Solution 1: Check if the map is nil before accessing
    if m == nil {
        fmt.Println("Map is nil")
    } else {
        value, ok := m["a"]
        if ok {
            fmt.Println(value)
        } else {
            fmt.Println("Key not found")
        }
    }
    
    // Solution 2: Initialize the map before use
    m = make(map[string]int)
    fmt.Println(m["a"]) // this will print 0
    m["a"] = 10
    fmt.Println(m["a"]) // this will print 10
} 
```
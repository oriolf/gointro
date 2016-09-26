package main

import "fmt"
import "runtime"

func main() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Println("Total: ", m.TotalAlloc / 1024, "KiB")
    fmt.Println("Used: ", m.Alloc / 1024, "KiB")
}
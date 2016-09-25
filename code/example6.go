package main
import "fmt"
import "time"
import "runtime"
func main() {
    PrintMem()
    for i := 0; i < 5; i++ {
        KeepMemory()
        time.Sleep(1 * time.Second)
        PrintMem()
    }
}

func PrintMem() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Println("Alloc: ", m.Alloc / 1024, "KiB")
}

func KeepMemory() {
    total := make([]byte, 200 * 1024 * 1024)
    
    fmt.Println("Total length: ", len(total))
    PrintMem()
}
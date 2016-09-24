package main
import "fmt"
import "time"
import "runtime"
func main() {
    go func() {
        t := time.NewTicker(5 * time.Second)
        var m *runtime.MemStats
        for range t.C {
            runtime.ReadMemStats(m)
            fmt.Println(*m)
        }
    }()

    time.Sleep(1 * time.Minute); KeepMemory(); time.Sleep(1 * time.Minute)
}

func KeepMemory() {
    total := make([][]byte, 12)
    for i := 0; i < 5; i++ {
        total = append(total, make([]byte, 50 * 1024 * 1024)); time.Sleep(5 * time.Second)
    }
}
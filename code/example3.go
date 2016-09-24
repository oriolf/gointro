package main

import "fmt"

type tree struct {
    v    int
    l, r *tree
}

func (t *tree) Sum() int {
    if t == nil {
        return 0
    }
    
    return t.v + t.l.Sum() + t.r.Sum()
}

func main() {
    var t *tree
    fmt.Println(t, t.Sum())
}
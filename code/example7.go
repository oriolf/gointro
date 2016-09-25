package main

import "github.com/nayarsystems/nxsugar-go"
import "github.com/jaracil/ei"

func main() {
    service := nxsugar.NewService("tcp://root:root@localhost:1717", "examples", nil)
    service.AddMethod("fib", Fib)
    service.Serve()
}

func Fib(task *nxsugar.Task) (interface{}, *nxsugar.JsonRpcErr) {
    v := ei.N(task.Params).M("v").IntZ()
    r := make([]int, 0)
    for i, j := 0, 1; j < v; i, j = i+j, i {
        r = append(r, i)
    }
    return r, nil
}
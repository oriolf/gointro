package main

import "fmt"
import "time"

type stateFn func() stateFn

func Ping() stateFn {
    fmt.Println("Ping!"); time.Sleep(2 * time.Second)
    return Pong
}

func Pong() stateFn {
    fmt.Println("Pong!"); time.Sleep(2 * time.Second)
    return Ping
}

func main() {
    for state := Ping; state != nil; {
        state = state()
    }
}
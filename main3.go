package main

import (
    "fmt"
)

// produce sends numbers from 1 to 10 on a channel
func produce(ch chan int) {
    for i := 1; i <= 10; i++ {
        ch <- i
    }
    close(ch)
}

// consume reads from the channel and prints the numbers
func consume(ch chan int) {
    for num := range ch {
        fmt.Println(num)
    }
}

func main() {
    ch := make(chan int)

    // Start produce in a new goroutine
    go produce(ch)

    // Start consume in a new goroutine
    go consume(ch)

    // Wait for both goroutines to finish
    fmt.Scanln() // to prevent main from exiting
}
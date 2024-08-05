package main

import (
    "fmt"
    "sync"
    "time"
)

// printNumbers prints numbers from 1 to 10
func printNumbers(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Millisecond * 10)
    }
}

// printLetters prints letters from 'a' to 'j'
func printLetters(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 'a'; i <= 'j'; i++ {
        fmt.Printf("%c\n", i)
        time.Sleep(time.Millisecond * 10)
    }
}

func main() {
    var wg sync.WaitGroup

    // Start printNumbers in a new goroutine
    wg.Add(1)
    go printNumbers(&wg)

    // Start printLetters in a new goroutine
    wg.Add(1)
    go printLetters(&wg)

    // Wait for both goroutines to finish
    wg.Wait()
}
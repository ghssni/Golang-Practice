package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// printNumbers prints numbers from 1 to 10
func printNumbers(wg *sync.WaitGroup, mutex *sync.Mutex, first bool) {
    defer wg.Done()
    if first {
        mutex.Lock()
        for i := 1; i <= 10; i++ {
            fmt.Println(i)
        }
        mutex.Unlock()
    } else {
        time.Sleep(time.Millisecond * 10)
        mutex.Lock()
        for i := 1; i <= 10; i++ {
            fmt.Println(i)
        }
        mutex.Unlock()
    }
}

// printLetters prints letters from 'a' to 'j'
func printLetters(wg *sync.WaitGroup, mutex *sync.Mutex, first bool) {
    defer wg.Done()
    if first {
        mutex.Lock()
        for i := 'a'; i <= 'j'; i++ {
            fmt.Printf("%c\n", i)
        }
        mutex.Unlock()
    } else {
        time.Sleep(time.Millisecond * 10)
        mutex.Lock()
        for i := 'a'; i <= 'j'; i++ {
            fmt.Printf("%c\n", i)
        }
        mutex.Unlock()
    }
}

func main() {
    var wg sync.WaitGroup
    var mutex sync.Mutex
    rand.Seed(time.Now().UnixNano())
    first := rand.Intn(2) == 0

    // Start printNumbers in a new goroutine
    wg.Add(1)
    go printNumbers(&wg, &mutex, first)

    // Start printLetters in a new goroutine
    wg.Add(1)
    go printLetters(&wg, &mutex, !first)

    // Wait for both goroutines to finish
    wg.Wait()
}
package main

import (
    "fmt"
    "time"
)

func sendNumbers(ch1 chan int, ch2 chan int) {
    for i := 1; i <= 20; i++ {
        if i%2 == 0 {
            ch1 <- i
        } else {
            ch2 <- i
        }
        time.Sleep(100 * time.Millisecond) // to slow down the sending process
    }
    close(ch1)
    close(ch2)
}

func main() {
    evenCh := make(chan int)
    oddCh := make(chan int)

    go sendNumbers(evenCh, oddCh)

    for evenCh != nil || oddCh != nil {
        select {
        case num, ok := <-evenCh:
            if !ok {
                evenCh = nil
                break
            }
            fmt.Println("Received even number:", num)
        case num, ok := <-oddCh:
            if !ok {
                oddCh = nil
                break
            }
            fmt.Println("Received odd number:", num)
        }
    }
}
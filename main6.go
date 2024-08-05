package main

import (
    "fmt"
    "time"
)

func sendNumbers(ch1 chan int, ch2 chan int, errCh chan error) {
    for i := 1; i <= 22; i++ {
        if i > 20 {
            errCh <- fmt.Errorf("error: %d is greater than 20", i)
        } else if i%2 == 0 {
            ch1 <- i
        } else {
            ch2 <- i
        }
        time.Sleep(100 * time.Millisecond) // to slow down the sending process
    }
    close(ch1)
    close(ch2)
    close(errCh)
}

func main() {
    evenCh := make(chan int)
    oddCh := make(chan int)
    errCh := make(chan error)

    go sendNumbers(evenCh, oddCh, errCh)

    for evenCh != nil || oddCh != nil || errCh != nil {
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
        case err, ok := <-errCh:
            if !ok {
                errCh = nil
                break
            }
            fmt.Println("Error:", err)
        }
    }
}
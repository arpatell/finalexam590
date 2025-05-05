package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func oddProducer(inCh chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 29; i += 2 {
        time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) // sleep 0–1500ms
        inCh <- i
    }
}

func evenProducer(inCh chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 2; i <= 30; i += 2 {
        time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
        inCh <- i
    }
}

func consumer(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for v := range inCh {
        time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // sleep 0–3000ms
        outCh <- v * v
    }
}

func filter(outCh <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    var prev int
    first := true
    for v := range outCh {
        if first || v > prev { // keep larger values after first
            fmt.Println(v)
            prev = v
            first = false
        }
    }
}

func main() {
    // bounded data buffers with capacity 5
    inCh := make(chan int, 5)
    outCh := make(chan int, 5)

    var prodWg sync.WaitGroup 
    var consWg sync.WaitGroup
    var filterWg sync.WaitGroup // waitgroups for each stage

    prodWg.Add(2)
    go oddProducer(inCh, &prodWg)
    go evenProducer(inCh, &prodWg)

    go func() {
        prodWg.Wait()
        close(inCh) // close after producers finish, stops deadlock between stages
    }()

    consWg.Add(2)
    go consumer(inCh, outCh, &consWg)
    go consumer(inCh, outCh, &consWg)

    go func() {
        consWg.Wait()
        close(outCh) // close after consumers finish
    }()

    filterWg.Add(1)
    go filter(outCh, &filterWg)
    filterWg.Wait()
}

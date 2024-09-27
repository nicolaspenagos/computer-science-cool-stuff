package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

func main() {
    n, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        fmt.Println(err)
    }

    start := time.Now()
    fmt.Printf("Classic Recursive Fibonacci :\nf(%d)=%d, executionTime:%d nanoseconds", n, slowFibonacci(n), time.Since(start).Nanoseconds())
    start = time.Now()
    fmt.Printf("\n\nMemoized Recursive Fibonacci :\nf(%d)=%d, executionTime:%d nanoseconds", n, optimizedFibonacci(n), time.Since(start).Nanoseconds())
    start = time.Now()
    fmt.Printf("\n\nIterative Fibonacci :\nf(%d)=%d, executionTime:%d nanoseconds", n, iterativeFibonacci(n), time.Since(start).Nanoseconds())
}

func optimizedFibonacci(n int64) int64 {
    fiboCache := make(map[int64]int64)
    return fibonacci(n, fiboCache)
}

func fibonacci(n int64, fiboCache map[int64]int64) int64 {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    i, exists := fiboCache[n-1]
    if !exists {
        i = fibonacci(n-1, fiboCache)
    }

    j, exists := fiboCache[n-2]
    if !exists {
        j = fibonacci(n-2, fiboCache)
    }

    fiboCache[n] = i + j
    return i + j
}

func slowFibonacci(n int64) int64 {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    return slowFibonacci(n-2) + slowFibonacci(n-1)
}

func iterativeFibonacci(n int64) int64 {
    prev := int64(0)
    prevPrev := int64(1)
    var fibonacci int64
    for i := int64(0); i < n; i++ {
        fibonacci = prev + prevPrev
        prevPrev = prev
        prev = fibonacci

    }
    return fibonacci
}
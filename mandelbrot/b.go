package main

import (
    "bytes"
    "fmt"
    "time"
)

const (
    BAILOUT        = 16
    MAX_ITERATIONS = 1000
)

func main() {
    var buf bytes.Buffer
    startTime := time.Now()
    for y := -39.0; y < 39.0; y++ {
        for x := -39.0; x < 39.0; x++ {
            if iterate(x/40.0, y/40.0) == 0.0 {
                buf.WriteRune('*')
            } else {
                buf.WriteRune(' ')
            }
        }
        buf.WriteRune('\n')
    }
    fmt.Print(buf.String())
    fmt.Printf("\nGo Elapsed %.3f\n", time.Now().Sub(startTime).Seconds())
}

func iterate(x, y float64) float64 {
    cr := y - 0.5
    ci := x
    zr := 0.0
    zi := 0.0
    i := 0.0
    for {
        i++
        temp := zr * zi
        zr2 := zr * zr
        zi2 := zi * zi
        zr = zr2 - zi2 + cr
        zi = temp + temp + ci
        if zi2+zr2 > BAILOUT {
            return i
        }
        if i > MAX_ITERATIONS {
            return 0.0
        }
    }
}

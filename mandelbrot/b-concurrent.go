package main

import (
    "bytes"
    "fmt"
    "strings"
    "time"
)

const (
    BAILOUT        = 16
    MAX_ITERATIONS = 1000
)

type line struct {
    number int
    text   string
}

func main() {
    out := make(chan *line, 80)
    startTime := time.Now()
    for y := -39.0; y < 39.0; y++ {
        // Fan out
        go lines(y, out)
    }
    // Fan in
    ls := make([]string, 80)
    for y := -39.0; y < 39.0; y++ {
        l := <-out
        ls[l.number+39] = l.text
    }
    fmt.Print(strings.Join(ls, ""))
    fmt.Printf("\nGo Elapsed %.3f\n", time.Now().Sub(startTime).Seconds())
}

func lines(y float64, out chan<- *line) {
    var buf bytes.Buffer
    for x := -39.0; x < 39.0; x++ {
        if iterate(x/40.0, y/40.0) == 0.0 {
            buf.WriteRune('*')
        } else {
            buf.WriteRune(' ')
        }
    }
    buf.WriteRune('\n')
    out <- &line{
        number: int(y),
        text:   buf.String(),
    }
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

